package gomake

import (
	"context"
	"time"
)

const (
	defaultReadTimeout  time.Duration = 1 * time.Second
	defaultWriteTimeout time.Duration = 1 * time.Second
	defaultIdleTimeout  time.Duration = 10 * time.Minute
)

type FileCTX interface {
	context.Context
	Context() (ctx context.Context, cancel context.CancelFunc)
}

type fileCTX struct {
	ctx context.Context
}

var tmpCTX context.Context = ctxParent

// IdleCTX returns a context with a timeout(deadline) based on
// the config.IdleTimeout setting.
//
// Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete:
//
//  func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//      ctx, cancel := parent.IdleCTX()
//      defer cancel()  // releases resources if slowOperation completes before timeout elapses
//      return slowOperation(ctx)
//  }
func (f fileCTX) IdleCTX() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(ctxEncrypt, config.IdleTimeout)
}

// Idle returns a context with a timeout(deadline) based on
// the config.ReadTimeout setting.
//
// Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete:
//
//  func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//      ctx, cancel := parent.IdleCTX()
//      defer cancel()  // releases resources if slowOperation completes before timeout elapses
//      return slowOperation(ctx)
//  }
func (f fileCTX) ReadCTX() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(ctxEncrypt, config.ReadTimeout)
}

// Idle returns a context with a timeout(deadline) based on
// the config.WriteTimeout setting.
//
// Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete:
//
//  func readOperationWithTimeout(parent context.Context, p []byte) (int, error) {
//      ctx, cancel := parent.IdleCTX()
//      defer cancel()  // releases resources if slowOperation completes before timeout elapses
//      return readOperation(ctx, p []byte)
//  }
func (f fileCTX) WriteCTX() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(ctxEncrypt, config.WriteTimeout)
}
