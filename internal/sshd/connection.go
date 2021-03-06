package sshd

import (
	"context"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/semaphore"

	"gitlab.com/gitlab-org/gitlab-shell/internal/metrics"

	"gitlab.com/gitlab-org/labkit/log"
)

type connection struct {
	begin              time.Time
	concurrentSessions *semaphore.Weighted
	remoteAddr         string
}

type channelHandler func(context.Context, ssh.Channel, <-chan *ssh.Request)

func newConnection(maxSessions int64, remoteAddr string) *connection {
	return &connection{
		begin:              time.Now(),
		concurrentSessions: semaphore.NewWeighted(maxSessions),
		remoteAddr:         remoteAddr,
	}
}

func (c *connection) handle(ctx context.Context, chans <-chan ssh.NewChannel, handler channelHandler) {
	ctxlog := log.WithContextFields(ctx, log.Fields{"remote_addr": c.remoteAddr})

	defer metrics.SshdConnectionDuration.Observe(time.Since(c.begin).Seconds())

	for newChannel := range chans {
		ctxlog.WithField("channel_type", newChannel.ChannelType()).Info("connection: handle: new channel requested")
		if newChannel.ChannelType() != "session" {
			ctxlog.Info("connection: handle: unknown channel type")
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		if !c.concurrentSessions.TryAcquire(1) {
			ctxlog.Info("connection: handle: too many concurrent sessions")
			newChannel.Reject(ssh.ResourceShortage, "too many concurrent sessions")
			metrics.SshdHitMaxSessions.Inc()
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			ctxlog.WithError(err).Error("connection: handle: accepting channel failed")
			c.concurrentSessions.Release(1)
			continue
		}

		go func() {
			defer c.concurrentSessions.Release(1)

			// Prevent a panic in a single session from taking out the whole server
			defer func() {
				if err := recover(); err != nil {
					ctxlog.WithField("recovered_error", err).Warn("panic handling session")
				}
			}()

			handler(ctx, channel, requests)
			ctxlog.Info("connection: handle: done")
		}()
	}
}
