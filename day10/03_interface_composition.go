package main

import "fmt"

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 3: Interface Composition
// ============================================================================
//
// COMPOSING INTERFACES
//
// Go encourages small, focused interfaces. You can combine them into
// larger interfaces using embedding - just like struct composition!
//
// Philosophy: "The bigger the interface, the weaker the abstraction."
//              - Rob Pike
//
// Small interfaces are:
//   - Easier to implement
//   - More reusable
//   - More flexible
//
// ============================================================================

// Small, focused interfaces
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string) error
}

type Closer interface {
	Close() error
}

// Composed interfaces - embedding combines them
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ============================================================================
// IMPLEMENTING COMPOSED INTERFACES
// ============================================================================

// File implements all three base interfaces, so it also implements
// ReadWriter and ReadWriteCloser automatically!
type File struct {
	Name    string
	Content string
	IsOpen  bool
}

func (f *File) Read() string {
	if !f.IsOpen {
		return "Error: file is closed"
	}
	return f.Content
}

func (f *File) Write(data string) error {
	if !f.IsOpen {
		return fmt.Errorf("cannot write: file is closed")
	}
	f.Content = data
	return nil
}

func (f *File) Close() error {
	if !f.IsOpen {
		return fmt.Errorf("file already closed")
	}
	f.IsOpen = false
	fmt.Printf("Closing file: %s\n", f.Name)
	return nil
}

// NetworkConnection implements ReadWriteCloser
type NetworkConnection struct {
	Address   string
	Connected bool
	Buffer    string
}

func (nc *NetworkConnection) Read() string {
	if !nc.Connected {
		return "Error: not connected"
	}
	return nc.Buffer
}

func (nc *NetworkConnection) Write(data string) error {
	if !nc.Connected {
		return fmt.Errorf("not connected to %s", nc.Address)
	}
	nc.Buffer = data
	fmt.Printf("Sent to %s: %s\n", nc.Address, data)
	return nil
}

func (nc *NetworkConnection) Close() error {
	if !nc.Connected {
		return fmt.Errorf("already disconnected")
	}
	nc.Connected = false
	fmt.Printf("Disconnected from %s\n", nc.Address)
	return nil
}

// ============================================================================
// FUNCTIONS THAT USE COMPOSED INTERFACES
// ============================================================================

// Copy only needs Reader and Writer
func Copy(dst Writer, src Reader) error {
	data := src.Read()
	return dst.Write(data)
}

// Process reads, processes, writes, then closes
func Process(rwc ReadWriteCloser, transform func(string) string) error {
	data := rwc.Read()
	processed := transform(data)
	if err := rwc.Write(processed); err != nil {
		return err
	}
	return rwc.Close()
}

// SafeClose works with anything that can be closed
func SafeClose(c Closer) {
	if err := c.Close(); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}
}

// ============================================================================
// STANDARD LIBRARY INTERFACE COMPOSITION
// ============================================================================
//
// The standard library uses this pattern extensively:
//
//   io.Reader     -> Read(p []byte) (n int, err error)
//   io.Writer     -> Write(p []byte) (n int, err error)
//   io.Closer     -> Close() error
//   io.ReadWriter -> Reader + Writer
//   io.ReadCloser -> Reader + Closer
//   io.WriteCloser -> Writer + Closer
//   io.ReadWriteCloser -> Reader + Writer + Closer
//
// ============================================================================

// ============================================================================
// INTERFACE EMBEDDING WITH ADDITIONAL METHODS
// ============================================================================

// Formatter interface
type Formatter interface {
	Format() string
}

// Validator interface
type Validator interface {
	Validate() error
}

// FormData composes interfaces AND adds methods
type FormData interface {
	Formatter
	Validator
	Submit() error
}

// RegistrationForm implements FormData
type RegistrationForm struct {
	Username string
	Email    string
	Password string
}

func (rf RegistrationForm) Format() string {
	return fmt.Sprintf("User: %s, Email: %s", rf.Username, rf.Email)
}

func (rf RegistrationForm) Validate() error {
	if rf.Username == "" {
		return fmt.Errorf("username is required")
	}
	if rf.Email == "" {
		return fmt.Errorf("email is required")
	}
	if len(rf.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

func (rf RegistrationForm) Submit() error {
	if err := rf.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	fmt.Println("Form submitted:", rf.Format())
	return nil
}

// Compile-time verification
var _ FormData = RegistrationForm{}

func main() {
	fmt.Println("=== Interface Composition ===")
	fmt.Println()

	// File implements all interfaces
	fmt.Println("--- File (ReadWriteCloser) ---")
	file := &File{Name: "document.txt", IsOpen: true}
	file.Write("Hello, World!")
	fmt.Println("Read:", file.Read())
	file.Close()
	fmt.Println("Read after close:", file.Read())

	fmt.Println()

	// NetworkConnection implements ReadWriteCloser
	fmt.Println("--- NetworkConnection ---")
	conn := &NetworkConnection{Address: "api.example.com:443", Connected: true}
	conn.Write("GET /data HTTP/1.1")
	conn.Buffer = "Response data here"
	fmt.Println("Received:", conn.Read())
	conn.Close()

	fmt.Println()

	// Copy function uses composed interface
	fmt.Println("--- Copy Function ---")
	src := &File{Name: "source.txt", Content: "Important data", IsOpen: true}
	dst := &File{Name: "dest.txt", IsOpen: true}
	Copy(dst, src)
	fmt.Println("Destination content:", dst.Read())

	fmt.Println()

	// Process function
	fmt.Println("--- Process Function ---")
	dataFile := &File{Name: "data.txt", Content: "hello world", IsOpen: true}
	toUpper := func(s string) string {
		result := ""
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				result += string(c - 32)
			} else {
				result += string(c)
			}
		}
		return result
	}
	Process(dataFile, toUpper)

	fmt.Println()

	// SafeClose with different types
	fmt.Println("--- SafeClose Function ---")
	f := &File{Name: "temp.txt", IsOpen: true}
	c := &NetworkConnection{Address: "localhost", Connected: true}
	SafeClose(f)
	SafeClose(c)
	// Try closing again - should show warning
	SafeClose(f)

	fmt.Println()

	// FormData with composed + extra methods
	fmt.Println("--- FormData (Composed + Extra) ---")
	form := RegistrationForm{
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "securepass123",
	}
	form.Submit()

	// Invalid form
	invalidForm := RegistrationForm{Username: "jane"}
	if err := invalidForm.Submit(); err != nil {
		fmt.Println("Error:", err)
	}
}

// ============================================================================
// TO RUN:
//   go run day10/03_interface_composition.go
//
// EXPECTED OUTPUT:
//   === Interface Composition ===
//
//   --- File (ReadWriteCloser) ---
//   Read: Hello, World!
//   Closing file: document.txt
//   Read after close: Error: file is closed
//
//   --- NetworkConnection ---
//   Sent to api.example.com:443: GET /data HTTP/1.1
//   Received: Response data here
//   Disconnected from api.example.com:443
//
//   --- Copy Function ---
//   Destination content: Important data
//
//   --- Process Function ---
//   Closing file: data.txt
//
//   --- SafeClose Function ---
//   Closing file: temp.txt
//   Disconnected from localhost
//   Warning: file already closed
//
//   --- FormData (Composed + Extra) ---
//   Form submitted: User: johndoe, Email: john@example.com
//   Error: validation failed: email is required
//
// EXERCISE:
//   1. Create a Database interface with Query() and Execute() methods
//   2. Create a Transactional interface that embeds Database and adds
//      BeginTransaction(), Commit(), and Rollback()
//   3. Implement a MockDatabase that implements Transactional
//   4. Write a function that accepts Database (not Transactional)
//
// KEY POINTS:
//   - Interfaces can embed other interfaces
//   - A type implementing the composed interface must implement ALL methods
//   - Prefer small, focused interfaces over large ones
//   - Standard library uses this pattern: io.ReadWriteCloser
//   - Functions should accept the smallest interface they need
//   - Interface composition enables flexible, reusable code
// ============================================================================
