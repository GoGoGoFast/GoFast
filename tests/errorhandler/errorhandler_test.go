package errorhandler_test

import (
	"GoAllInOne/pkg/errorhandler"
	"errors"
	"os"
	"strings"
	"testing"
)

// 自定义错误处理器用于测试
type TestErrorHandler struct {
	handledErrors []error
}

func (h *TestErrorHandler) HandleError(err error) {
	h.handledErrors = append(h.handledErrors, err)
}

func TestErrorHandlerPackage(t *testing.T) {
	// 测试错误类型定义
	if errorhandler.ErrNotFound.Error() != "not found" {
		t.Errorf("expected 'not found', got '%s'", errorhandler.ErrNotFound.Error())
	}

	// 测试包装和解包错误
	originalErr := errors.New("original error")
	wrappedErr := errorhandler.Wrap(originalErr, "additional context")
	if !errors.Is(wrappedErr, originalErr) {
		t.Errorf("expected wrapped error to contain original error")
	}
	if unwrappedErr := errorhandler.Unwrap(wrappedErr); unwrappedErr != originalErr {
		t.Errorf("expected unwrapped error to be original error")
	}

	// 测试自定义错误创建
	customErr := errorhandler.NewError(1001, "Resource not found", errorhandler.Error, map[string]interface{}{"resource": "file.txt"}, originalErr)
	if customErr.Code != 1001 || customErr.Message != "Resource not found" || customErr.Level != errorhandler.Error {
		t.Errorf("custom error fields do not match expected values")
	}

	// 测试日志记录（控制台）
	errorhandler.LogError(customErr)

	// 测试日志记录（文件）
	logFilePath := "test_errors.log"
	defer os.Remove(logFilePath)
	if err := errorhandler.InitLogFile(logFilePath); err != nil {
		t.Fatalf("failed to initialize log file: %v", err)
	}
	errorhandler.LogError(customErr)

	// 检查日志文件内容
	logFileContent, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	if !strings.Contains(string(logFileContent), customErr.Error()) {
		t.Errorf("log file does not contain expected error message")
	}

	// 测试错误通知（仅打印输出）
	errorhandler.NotifyError(customErr, "admin@example.com")

	// 测试自定义错误处理器注册和触发
	testHandler := &TestErrorHandler{}
	errorhandler.RegisterErrorHandler(testHandler)
	errorhandler.TriggerCustomErrorHandlers(customErr)
	if len(testHandler.handledErrors) != 1 || testHandler.handledErrors[0] != customErr {
		t.Errorf("custom handler did not receive the expected error")
	}

	// 测试错误聚合
	aggErr := errorhandler.NewAggregateError([]error{customErr, errors.New("another sample error")})
	if len(aggErr.Errors) != 2 {
		t.Errorf("expected 2 errors in aggregate error, got %d", len(aggErr.Errors))
	}

	// 测试多语言支持
	localizedMessage := errorhandler.GetLocalizedMessage(1001, "zh")
	if localizedMessage != "资源未找到" {
		t.Errorf("expected '资源未找到', got '%s'", localizedMessage)
	}
}
