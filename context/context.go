package context

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"go.uber.org/zap"
	"maschine.io/core"
	"maschine.io/core/params"
)

const (
	INPUTKEY       = "input"
	LOGGERKEY      = "log"
	CREDENTIALSKEY = "credentials"
	PARAMSKEY      = "params"
)

type Context struct {
	// This mutex protects Keys map.
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context.
	Keys map[string]any

	done     chan struct{}
	err      error
	deadline time.Time
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes c.Keys if it was not used previously.
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
	c.mu.Unlock()
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	value, exists = c.Keys[key]
	c.mu.RUnlock()
	return
}

// GetCredential returns the credential by key if key exists
func (c *Context) GetCredential(key string) (value any, exists bool) {
	c.mu.RLock()
	if cred, credExists := c.Keys[CREDENTIALSKEY]; credExists {
		if c, ok := cred.(map[string]any); ok {
			value, exists = c[key], true
		}
	}
	c.mu.RUnlock()
	return
}

func (c *Context) CredentialsExists() bool {
	c.mu.RLock()
	_, credExists := c.Keys[CREDENTIALSKEY]
	c.mu.RUnlock()
	return credExists
}

// Background returns a new context with a background value
func Background() *Context {
	return &Context{
		Keys: map[string]any{},
		done: make(chan struct{}),
	}
}

// Deadline returns the time when work done on behalf of this context
// should be canceled. Deadline returns ok==false when no deadline is set.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.deadline.IsZero() {
		return time.Time{}, false
	}
	return c.deadline, true
}

// Done returns a channel that's closed when work done on behalf of this
// context should be canceled. Done may return nil if this context can
// never be canceled.
func (c *Context) Done() <-chan struct{} {
	return c.done
}

// Err returns a non-nil error value after Done is closed. Err returns
// Canceled if the context was canceled or DeadlineExceeded if the
// context's deadline passed. No other values for Err are defined.
func (c *Context) Err() error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.err
}

// GetInputAsInterface returns the input as interface{}
func (c *Context) GetInputAsInterface() (any, error) {
	value := c.GetInput()
	return core.InputToBytes(value)
}

func (c *Context) GetInputAsMap() (result map[string]any, err error) {
	resultMap := map[string]any{}
	value := c.GetInput()
	if err = json.Unmarshal(value, &resultMap); err != nil {
		return
	}
	return resultMap, nil
}

// GetInput returns the input if it is present in this context
// If no input is given it return a empty json object as []byte
func (c *Context) GetInput() (value []byte) {
	var input []byte = []byte("{}")
	gotInput, exists := c.Get(INPUTKEY)
	if !exists {
		return input
	}
	if inputFromContext, ok := gotInput.([]byte); ok {
		return inputFromContext
	}
	return input
}

func (c *Context) SetInput(input []byte) {
	c.Set(INPUTKEY, input)
}

func (c *Context) SetParams(params *params.Parameter) {
	c.Set(PARAMSKEY, params)
}

func (c *Context) GetParams() *params.Parameter {
	if p, ok := c.Get(PARAMSKEY); ok {
		return p.(*params.Parameter)
	}
	return nil
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (c *Context) MustGet(key string) any {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetStringWithDefault(key string, def string) (s string) {
	s = c.GetString(key)
	if len(s) == 0 {
		return def
	}
	return
}

// GetBytes returns the value associated with the key as a []byte slice.
func (c *Context) GetBytes(key string) (b []byte) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.([]byte)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetUint returns the value associated with the key as an unsigned integer.
func (c *Context) GetUint(key string) (ui uint) {
	if val, ok := c.Get(key); ok && val != nil {
		ui, _ = val.(uint)
	}
	return
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (c *Context) GetUint64(key string) (ui64 uint64) {
	if val, ok := c.Get(key); ok && val != nil {
		ui64, _ = val.(uint64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key string) (sm map[string]any) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]any)
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

/************************************/
/***** Parameters               *****/
/************************************/

// func (c *Context) SetParams(params *[]Parameter)

func (c *Context) SetLogger(logger *zap.Logger) {
	c.Set(LOGGERKEY, logger)
}

func (c *Context) DefaultLogger() *zap.Logger {
	return c.GetLogger(LOGGERKEY)
}

// GetLogger returns the logger from context; If no logger exists a new production logger will be returned
func (c *Context) GetLogger(lg string) (logger *zap.Logger) {
	var check bool
	if l, ok := c.Get(lg); ok {
		if logger, check = l.(*zap.Logger); check {
			return logger
		}
	}
	logger, _ = zap.NewProduction()
	return
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(key any) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if keyAsString, ok := key.(string); ok {
		if val, exists := c.Get(keyAsString); exists {
			return val
		}
	}
	return nil
}

// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
func WithCancel(parent *Context) (*Context, func()) {
	ctx := &Context{
		Keys: parent.Keys,
		done: make(chan struct{}),
	}

	cancel := func() {
		ctx.mu.Lock()
		defer ctx.mu.Unlock()
		if ctx.err == nil {
			ctx.err = context.Canceled
			close(ctx.done)
		}
	}

	go func() {
		select {
		case <-parent.Done():
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is closed,
// whichever happens first.
func WithDeadline(parent *Context, d time.Time) (*Context, func()) {
	ctx := &Context{
		Keys:     parent.Keys,
		done:     make(chan struct{}),
		deadline: d,
	}

	cancel := func() {
		ctx.mu.Lock()
		defer ctx.mu.Unlock()
		if ctx.err == nil {
			ctx.err = context.DeadlineExceeded
			close(ctx.done)
		}
	}

	go func() {
		select {
		case <-parent.Done():
			cancel()
		case <-time.After(time.Until(d)):
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
func WithTimeout(parent *Context, timeout time.Duration) (*Context, func()) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent *Context, key, val any) *Context {
	ctx := &Context{
		Keys:     parent.Keys,
		done:     parent.done,
		err:      parent.err,
		deadline: parent.deadline,
	}
	if keyAsString, ok := key.(string); ok {
		ctx.Set(keyAsString, val)
	}
	return ctx
}
