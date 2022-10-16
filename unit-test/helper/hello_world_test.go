package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Romzi")
	}
}

func BenchmarkHelloWorldFarhan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Farhan")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Romzi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Romzi")
		}
	})
	b.Run("Farhan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Farhan")
		}
	})
}

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "Romzi",
			request: "Romzi",
		},
		{
			name:    "Farhan",
			request: "Farhan",
		}}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Romzi",
			request:  "Romzi",
			expected: "Hello Romzi",
		},
		{
			name:     "Farhan",
			request:  "Farhan",
			expected: "Hello Farhan",
		},
		{
			name:     "Ghozi",
			request:  "Ghozi",
			expected: "Hello Ghozi",
		}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Romzi")
	if result != "Hello Romzi" {
		// t.Fail()
		t.Error("Harusnya Hi Romzi")
	}

	fmt.Println("Test Hello World")
}

func TestHelloDunia(t *testing.T) {
	result := HelloWorld("Dunia")
	if result != "Hello Dunia" {
		// t.FailNow()
		t.Fatal("Harusnya Hi Dunia")
	}

	fmt.Println("Test Hello Dunia")
}

func TestAssertion(t *testing.T) {
	result := HelloWorld("Romzi")
	assert.Equal(t, "Hello Romzi", result, "Result mush be 'Hello Romzi'")
	fmt.Println("Test using Assertion")
}

func TestRequire(t *testing.T) {
	result := HelloWorld("Romzi")
	require.Equal(t, "Hello Romzi", result, "Result mush be 'Hello Romzi'")
	fmt.Println("Test using Assertion")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

	result := HelloWorld("Romzi")
	require.Equal(t, "Hello Romzi", result, "Result mush be 'Hello Romzi'")
	fmt.Println("Test using Assertion")
}

func TestSubTest(t *testing.T) {
	t.Run("Romzi", func(t *testing.T) {
		result := HelloWorld("Romzi")
		require.Equal(t, "Hello Romzi", result, "Result mush be 'Hello Romzi'")
	})

	t.Run("Farhan", func(t *testing.T) {
		result := HelloWorld("Farhan")
		require.Equal(t, "Hello Farhan", result, "Result mush be 'Hello Farhan'")
	})
}
