package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ini adalah Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Sehan",
			request: "Sehan",
		},
		{
			name:    "Givi",
			request: "Givi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// ini adalah Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sehan")
	}
}

func BenchmarkHelloWorldGivi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Givi")
	}
}

// Ini adalah Table Test
func TestTabelHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Sehan",
			request:  "Sehan",
			expected: "Hello Sehan",
		},
		{
			name:     "Givi",
			request:  "Givi",
			expected: "Hello Givi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// ini adalah Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Sehan", func(t *testing.T) {
		result := HelloWorld("Sehan")
		require.Equal(t, "Hello Sehan", result)
	})
	t.Run("Givi", func(t *testing.T) {
		result := HelloWorld("Givi")
		require.Equal(t, "Hello Givi", result)
	})
}

// ini adalah Before and After
func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run() // eksekusi semua unit test

	fmt.Println("After Unit Test")
}

// ini adalah Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Sehan")
	require.Equal(t, "Hello Sehan", result, "Result must be 'Hello Sehan'")
}

// ini adalah Require.
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sehan")
	require.Equal(t, "Hello Sehan", result, "Result must be 'Hello Sehan'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// Ini adalah Assertion.
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Sehan")
	assert.Equal(t, "Hello Sehan", result, "Result must be 'Hello Sehan'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// ini adalah contoh Unit test dan Menggagalkan test.
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sehan")

	if result != "Hello Sehan" {
		//Error
		t.Error("Result is no Hello Sehan")
	}

	fmt.Println("TestHelloWorldSehan Done")
}

func TestHelloWorldGivi(t *testing.T) {
	result := HelloWorld("Givi")

	if result != "Hello Givi" {
		//unit test failed
		t.Fatal("Result is no Hello Givi")
	}

	fmt.Println("TesHelloWordlGivi Done")
}
