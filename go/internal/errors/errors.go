package errors

import (
	stderrors "errors"
	"fmt"
	"runtime"
	"strings"
)

// Error は階層的なエラー情報とスタックトレースを保持します。
type Error struct {
	msg   string    // エラーメッセージ
	cause error     // ラップしている元のエラー
	stack []uintptr // キャプチャしたスタックトレース
}

// New は新たな HierError を作成し、スタックトレースをキャプチャします。
func New(msg string) error {
	return &Error{
		msg:   msg,
		stack: captureStack(),
	}
}

// Wrap は既存のエラーに新たなメッセージを付与してラップします。
// err が nil の場合は nil を返します。
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &Error{
		msg:   msg,
		cause: err,
		stack: captureStack(),
	}
}

// captureStack は現在のスタックトレースをキャプチャします。
// 呼び出し元の不要なフレーム（captureStack自身など）はスキップします。
func captureStack() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	// 3フレーム分スキップ（captureStack、自身のラッパー、さらにラッパーの呼び出し元）
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

// Error は error インターフェースの実装です。
// 内部に原因エラーがある場合は階層的にメッセージを連結して返します。
func (e *Error) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	}
	return e.msg
}

// Unwrap は内部に保持しているエラー（原因）を返します。
// これにより、標準の errors.Unwrap, errors.Is, errors.As が利用可能となります。
func (e *Error) Unwrap() error {
	return e.cause
}

// FormatStack はキャプチャしたスタックトレースを整形して文字列として返します。
func (e *Error) FormatStack() string {
	var sb strings.Builder
	frames := runtime.CallersFrames(e.stack)
	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return sb.String()
}

// Format は fmt.Formatter インターフェースを実装し、
// "%+v" で詳細なエラー情報（メッセージ＋スタックトレース）が出力できるようにします。
func (e *Error) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Flag('+') {
			// 詳細なフォーマット：エラーメッセージとスタックトレースを表示
			fmt.Fprintf(f, "Error: %s\nStack Trace:\n%s", e.Error(), indentLines(e.FormatStack(), "  "))
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(f, e.Error())
	case 'q':
		fmt.Fprintf(f, "%q", e.Error())
	}
}

// FormatErrorWithStack はエラーの階層情報および各階層のスタックトレースを整形して返します。
func FormatErrorWithStack(err error) string {
	var sb strings.Builder
	formatErrorRecursive(err, &sb, 0)
	return sb.String()
}

// formatErrorRecursive はエラーのチェーンを再帰的に整形して出力します。
func formatErrorRecursive(err error, sb *strings.Builder, level int) {
	indent := strings.Repeat("  ", level)
	sb.WriteString(fmt.Sprintf("%sError: %s\n", indent, err.Error()))

	// HierError 型のエラーであればスタックトレースを出力
	var he *Error
	if stderrors.As(err, &he) {
		sb.WriteString(fmt.Sprintf("%sStack Trace:\n%s\n", indent, indentLines(he.FormatStack(), indent+"  ")))
	}

	// 内部の原因エラーがあれば再帰的に出力
	if cause := stderrors.Unwrap(err); cause != nil {
		sb.WriteString(fmt.Sprintf("%sCaused by:\n", indent))
		formatErrorRecursive(cause, sb, level+1)
	}
}

// indentLines は各行の先頭に指定のインデント文字列を追加します。
func indentLines(s, indent string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}
	return strings.Join(lines, "\n")
}
