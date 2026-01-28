package main

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 6: Challenge - Plugin-Based Document Processing System
// ============================================================================
//
// CHALLENGE: Build a document processing system using interfaces
//
// You'll create a flexible document processor that can:
//   1. Read documents from different sources (files, URLs, databases)
//   2. Transform documents (convert format, compress, encrypt)
//   3. Write documents to different destinations
//   4. Support plugins for extensibility
//
// This challenge demonstrates:
//   - Interface design for flexibility
//   - Interface composition
//   - Type assertions for optional features
//   - The strategy pattern using interfaces
//   - Plugin architecture
//
// ============================================================================

// ============================================================================
// CORE INTERFACES
// ============================================================================

// Document represents any document in the system
type Document interface {
	GetContent() string
	GetMetadata() Metadata
}

// Metadata holds document information
type Metadata struct {
	Name      string
	Size      int
	Created   time.Time
	Modified  time.Time
	MimeType  string
	Author    string
	Tags      []string
}

// DocumentReader can read documents from a source
type DocumentReader interface {
	Read(source string) (Document, error)
	SupportedSources() []string
}

// DocumentWriter can write documents to a destination
type DocumentWriter interface {
	Write(doc Document, destination string) error
	SupportedDestinations() []string
}

// DocumentTransformer can transform documents
type DocumentTransformer interface {
	Transform(doc Document) (Document, error)
	Name() string
}

// ============================================================================
// COMPOSED INTERFACES
// ============================================================================

// DocumentProcessor combines reading and writing
type DocumentProcessor interface {
	DocumentReader
	DocumentWriter
}

// Plugin represents a loadable plugin
type Plugin interface {
	Name() string
	Version() string
	Initialize() error
}

// TransformerPlugin is a plugin that provides transformations
type TransformerPlugin interface {
	Plugin
	GetTransformers() []DocumentTransformer
}

// ============================================================================
// OPTIONAL CAPABILITY INTERFACES
// ============================================================================

// Searchable documents can be searched
type Searchable interface {
	Search(query string) []SearchResult
}

// SearchResult represents a search match
type SearchResult struct {
	Line    int
	Column  int
	Context string
}

// Versionable documents support versioning
type Versionable interface {
	GetVersion() int
	GetHistory() []VersionInfo
}

// VersionInfo represents a version entry
type VersionInfo struct {
	Version   int
	Timestamp time.Time
	Author    string
	Comment   string
}

// Encryptable documents can be encrypted/decrypted
type Encryptable interface {
	Encrypt(key string) error
	Decrypt(key string) error
	IsEncrypted() bool
}

// ============================================================================
// CONCRETE IMPLEMENTATIONS
// ============================================================================

// TextDocument is a basic document implementation
type TextDocument struct {
	content   string
	metadata  Metadata
	version   int
	history   []VersionInfo
	encrypted bool
}

// NewTextDocument creates a new text document
func NewTextDocument(name, content string, author string) *TextDocument {
	now := time.Now()
	return &TextDocument{
		content: content,
		metadata: Metadata{
			Name:     name,
			Size:     len(content),
			Created:  now,
			Modified: now,
			MimeType: "text/plain",
			Author:   author,
			Tags:     []string{},
		},
		version: 1,
		history: []VersionInfo{
			{Version: 1, Timestamp: now, Author: author, Comment: "Initial creation"},
		},
	}
}

func (t *TextDocument) GetContent() string {
	return t.content
}

func (t *TextDocument) GetMetadata() Metadata {
	return t.metadata
}

// TextDocument implements Searchable
func (t *TextDocument) Search(query string) []SearchResult {
	var results []SearchResult
	lines := strings.Split(t.content, "\n")

	for i, line := range lines {
		if idx := strings.Index(strings.ToLower(line), strings.ToLower(query)); idx != -1 {
			results = append(results, SearchResult{
				Line:    i + 1,
				Column:  idx + 1,
				Context: line,
			})
		}
	}
	return results
}

// TextDocument implements Versionable
func (t *TextDocument) GetVersion() int {
	return t.version
}

func (t *TextDocument) GetHistory() []VersionInfo {
	return t.history
}

func (t *TextDocument) UpdateContent(content, author, comment string) {
	t.content = content
	t.version++
	t.metadata.Modified = time.Now()
	t.metadata.Size = len(content)
	t.history = append(t.history, VersionInfo{
		Version:   t.version,
		Timestamp: time.Now(),
		Author:    author,
		Comment:   comment,
	})
}

// TextDocument implements Encryptable (simulated)
func (t *TextDocument) Encrypt(key string) error {
	if t.encrypted {
		return fmt.Errorf("document already encrypted")
	}
	// Simulated encryption - just reverse the content
	runes := []rune(t.content)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	t.content = string(runes)
	t.encrypted = true
	return nil
}

func (t *TextDocument) Decrypt(key string) error {
	if !t.encrypted {
		return fmt.Errorf("document is not encrypted")
	}
	// Reverse again to decrypt
	runes := []rune(t.content)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	t.content = string(runes)
	t.encrypted = false
	return nil
}

func (t *TextDocument) IsEncrypted() bool {
	return t.encrypted
}

// Compile-time interface verification
var _ Document = (*TextDocument)(nil)
var _ Searchable = (*TextDocument)(nil)
var _ Versionable = (*TextDocument)(nil)
var _ Encryptable = (*TextDocument)(nil)

// ============================================================================
// DOCUMENT READERS
// ============================================================================

// FileReader reads documents from the filesystem
type FileReader struct{}

func (fr FileReader) Read(source string) (Document, error) {
	// Simulated file reading
	content := fmt.Sprintf("Content loaded from file: %s\nThis is simulated file content.\nLine 3 of the document.", source)
	return NewTextDocument(source, content, "FileSystem"), nil
}

func (fr FileReader) SupportedSources() []string {
	return []string{"file://", "local path"}
}

// URLReader reads documents from URLs
type URLReader struct{}

func (ur URLReader) Read(source string) (Document, error) {
	// Simulated URL reading
	content := fmt.Sprintf("Content fetched from URL: %s\n<html><body>Web content here</body></html>", source)
	doc := NewTextDocument(source, content, "WebFetcher")
	doc.metadata.MimeType = "text/html"
	return doc, nil
}

func (ur URLReader) SupportedSources() []string {
	return []string{"http://", "https://"}
}

// DatabaseReader reads documents from database
type DatabaseReader struct {
	ConnectionString string
}

func (dr DatabaseReader) Read(source string) (Document, error) {
	// Simulated database reading
	content := fmt.Sprintf("Document retrieved from database\nTable: %s\nRecord data here...", source)
	return NewTextDocument(source, content, "Database"), nil
}

func (dr DatabaseReader) SupportedSources() []string {
	return []string{"db://", "sql://"}
}

// ============================================================================
// DOCUMENT WRITERS
// ============================================================================

// FileWriter writes documents to filesystem
type FileWriter struct{}

func (fw FileWriter) Write(doc Document, destination string) error {
	meta := doc.GetMetadata()
	fmt.Printf("Writing to file: %s\n", destination)
	fmt.Printf("  Content length: %d bytes\n", len(doc.GetContent()))
	fmt.Printf("  MIME type: %s\n", meta.MimeType)
	return nil
}

func (fw FileWriter) SupportedDestinations() []string {
	return []string{"file://", "local path"}
}

// ConsoleWriter outputs documents to console
type ConsoleWriter struct {
	Verbose bool
}

func (cw ConsoleWriter) Write(doc Document, destination string) error {
	meta := doc.GetMetadata()
	fmt.Println("=== Console Output ===")
	if cw.Verbose {
		fmt.Printf("Name: %s\n", meta.Name)
		fmt.Printf("Author: %s\n", meta.Author)
		fmt.Printf("Type: %s\n", meta.MimeType)
		fmt.Println("---")
	}
	fmt.Println(doc.GetContent())
	fmt.Println("======================")
	return nil
}

func (cw ConsoleWriter) SupportedDestinations() []string {
	return []string{"console://", "stdout"}
}

// ============================================================================
// DOCUMENT TRANSFORMERS
// ============================================================================

// UppercaseTransformer converts content to uppercase
type UppercaseTransformer struct{}

func (ut UppercaseTransformer) Transform(doc Document) (Document, error) {
	content := strings.ToUpper(doc.GetContent())
	meta := doc.GetMetadata()
	return NewTextDocument(meta.Name, content, meta.Author), nil
}

func (ut UppercaseTransformer) Name() string {
	return "uppercase"
}

// TrimTransformer trims whitespace
type TrimTransformer struct{}

func (tt TrimTransformer) Transform(doc Document) (Document, error) {
	lines := strings.Split(doc.GetContent(), "\n")
	var trimmed []string
	for _, line := range lines {
		trimmed = append(trimmed, strings.TrimSpace(line))
	}
	content := strings.Join(trimmed, "\n")
	meta := doc.GetMetadata()
	return NewTextDocument(meta.Name, content, meta.Author), nil
}

func (tt TrimTransformer) Name() string {
	return "trim"
}

// LineNumberTransformer adds line numbers
type LineNumberTransformer struct{}

func (ln LineNumberTransformer) Transform(doc Document) (Document, error) {
	lines := strings.Split(doc.GetContent(), "\n")
	var numbered []string
	for i, line := range lines {
		numbered = append(numbered, fmt.Sprintf("%3d: %s", i+1, line))
	}
	content := strings.Join(numbered, "\n")
	meta := doc.GetMetadata()
	return NewTextDocument(meta.Name+"_numbered", content, meta.Author), nil
}

func (ln LineNumberTransformer) Name() string {
	return "line-numbers"
}

// ============================================================================
// PLUGIN SYSTEM
// ============================================================================

// TextToolsPlugin provides text transformation tools
type TextToolsPlugin struct {
	name    string
	version string
}

func NewTextToolsPlugin() *TextToolsPlugin {
	return &TextToolsPlugin{
		name:    "TextTools",
		version: "1.0.0",
	}
}

func (p *TextToolsPlugin) Name() string {
	return p.name
}

func (p *TextToolsPlugin) Version() string {
	return p.version
}

func (p *TextToolsPlugin) Initialize() error {
	fmt.Printf("Initializing plugin: %s v%s\n", p.name, p.version)
	return nil
}

func (p *TextToolsPlugin) GetTransformers() []DocumentTransformer {
	return []DocumentTransformer{
		UppercaseTransformer{},
		TrimTransformer{},
		LineNumberTransformer{},
	}
}

// Verify plugin interfaces
var _ Plugin = (*TextToolsPlugin)(nil)
var _ TransformerPlugin = (*TextToolsPlugin)(nil)

// ============================================================================
// DOCUMENT PROCESSING ENGINE
// ============================================================================

// ProcessingEngine orchestrates document processing
type ProcessingEngine struct {
	readers      map[string]DocumentReader
	writers      map[string]DocumentWriter
	transformers []DocumentTransformer
	plugins      []Plugin
}

func NewProcessingEngine() *ProcessingEngine {
	return &ProcessingEngine{
		readers:      make(map[string]DocumentReader),
		writers:      make(map[string]DocumentWriter),
		transformers: []DocumentTransformer{},
		plugins:      []Plugin{},
	}
}

func (e *ProcessingEngine) RegisterReader(name string, reader DocumentReader) {
	e.readers[name] = reader
}

func (e *ProcessingEngine) RegisterWriter(name string, writer DocumentWriter) {
	e.writers[name] = writer
}

func (e *ProcessingEngine) RegisterTransformer(transformer DocumentTransformer) {
	e.transformers = append(e.transformers, transformer)
}

func (e *ProcessingEngine) LoadPlugin(plugin Plugin) error {
	if err := plugin.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize plugin %s: %w", plugin.Name(), err)
	}

	e.plugins = append(e.plugins, plugin)

	// Check if plugin provides transformers
	if tp, ok := plugin.(TransformerPlugin); ok {
		for _, t := range tp.GetTransformers() {
			e.RegisterTransformer(t)
			fmt.Printf("  Registered transformer: %s\n", t.Name())
		}
	}

	return nil
}

func (e *ProcessingEngine) Process(readerName, source string, transformerNames []string, writerName, destination string) error {
	// Get reader
	reader, ok := e.readers[readerName]
	if !ok {
		return fmt.Errorf("reader not found: %s", readerName)
	}

	// Read document
	doc, err := reader.Read(source)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	fmt.Printf("Read document: %s (%d bytes)\n", doc.GetMetadata().Name, len(doc.GetContent()))

	// Apply transformers
	for _, name := range transformerNames {
		for _, t := range e.transformers {
			if t.Name() == name {
				doc, err = t.Transform(doc)
				if err != nil {
					return fmt.Errorf("transform error (%s): %w", name, err)
				}
				fmt.Printf("Applied transformer: %s\n", name)
				break
			}
		}
	}

	// Get writer
	writer, ok := e.writers[writerName]
	if !ok {
		return fmt.Errorf("writer not found: %s", writerName)
	}

	// Write document
	if err := writer.Write(doc, destination); err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	return nil
}

// ProcessWithCapabilities checks for optional document capabilities
func (e *ProcessingEngine) ProcessWithCapabilities(doc Document) {
	fmt.Println("\n--- Document Capabilities ---")
	meta := doc.GetMetadata()
	fmt.Printf("Document: %s\n", meta.Name)

	// Check for Searchable
	if searchable, ok := doc.(Searchable); ok {
		fmt.Println("  [x] Searchable")
		results := searchable.Search("content")
		if len(results) > 0 {
			fmt.Printf("      Found %d matches for 'content'\n", len(results))
		}
	} else {
		fmt.Println("  [ ] Searchable")
	}

	// Check for Versionable
	if versionable, ok := doc.(Versionable); ok {
		fmt.Println("  [x] Versionable")
		fmt.Printf("      Current version: %d\n", versionable.GetVersion())
		fmt.Printf("      History entries: %d\n", len(versionable.GetHistory()))
	} else {
		fmt.Println("  [ ] Versionable")
	}

	// Check for Encryptable
	if encryptable, ok := doc.(Encryptable); ok {
		fmt.Println("  [x] Encryptable")
		fmt.Printf("      Currently encrypted: %v\n", encryptable.IsEncrypted())
	} else {
		fmt.Println("  [ ] Encryptable")
	}
}

func main() {
	fmt.Println("=== Document Processing System ===")
	fmt.Println()

	// Create processing engine
	engine := NewProcessingEngine()

	// Register readers
	engine.RegisterReader("file", FileReader{})
	engine.RegisterReader("url", URLReader{})
	engine.RegisterReader("db", DatabaseReader{ConnectionString: "localhost:5432"})
	fmt.Println("Registered readers: file, url, db")

	// Register writers
	engine.RegisterWriter("file", FileWriter{})
	engine.RegisterWriter("console", ConsoleWriter{Verbose: true})
	fmt.Println("Registered writers: file, console")

	fmt.Println()

	// Load plugin
	fmt.Println("--- Loading Plugins ---")
	plugin := NewTextToolsPlugin()
	engine.LoadPlugin(plugin)

	fmt.Println()

	// Process a document
	fmt.Println("--- Processing Document ---")
	err := engine.Process(
		"file", "report.txt",
		[]string{"trim", "line-numbers"},
		"console", "stdout",
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()

	// Demonstrate capability checking
	fmt.Println("--- Checking Document Capabilities ---")
	doc := NewTextDocument(
		"analysis.txt",
		"This document has content.\nIt supports multiple features.\nSearch for content here.",
		"System",
	)

	// Update the document to show versioning
	doc.UpdateContent(
		"Updated content.\nNew version with changes.\nSearch still works.",
		"Admin",
		"Updated for demo",
	)

	engine.ProcessWithCapabilities(doc)

	fmt.Println()

	// Demonstrate encryption
	fmt.Println("--- Encryption Demo ---")
	secret := NewTextDocument("secret.txt", "Top secret information!", "Agent")
	fmt.Println("Original:", secret.GetContent())

	secret.Encrypt("password123")
	fmt.Println("Encrypted:", secret.GetContent())
	fmt.Println("Is encrypted:", secret.IsEncrypted())

	secret.Decrypt("password123")
	fmt.Println("Decrypted:", secret.GetContent())

	fmt.Println()

	// Show search capability
	fmt.Println("--- Search Demo ---")
	searchDoc := NewTextDocument("article.txt",
		"Go interfaces are powerful.\nThey enable polymorphism.\nInterfaces in Go are implicit.",
		"Author")

	results := searchDoc.Search("interfaces")
	fmt.Printf("Found %d matches for 'interfaces':\n", len(results))
	for _, r := range results {
		fmt.Printf("  Line %d, Col %d: %s\n", r.Line, r.Column, r.Context)
	}
}

// ============================================================================
// TO RUN:
//   go run day10/06_challenge.go
//
// EXPECTED OUTPUT:
//   === Document Processing System ===
//
//   Registered readers: file, url, db
//   Registered writers: file, console
//
//   --- Loading Plugins ---
//   Initializing plugin: TextTools v1.0.0
//     Registered transformer: uppercase
//     Registered transformer: trim
//     Registered transformer: line-numbers
//
//   --- Processing Document ---
//   Read document: report.txt (93 bytes)
//   Applied transformer: trim
//   Applied transformer: line-numbers
//   === Console Output ===
//   Name: report.txt_numbered
//   ... (more output)
//
// EXERCISES:
//   1. Add a MarkdownDocument type that implements Document and has
//      a method ToHTML() string
//   2. Create a CompressionTransformer that simulates compression
//   3. Implement a CloudWriter that "uploads" to cloud storage
//   4. Add a StatisticsPlugin that provides a WordCountTransformer
//   5. Create a pipeline builder: engine.Pipeline().Read("file", "x.txt")
//                                       .Transform("uppercase")
//                                       .Write("console")
//                                       .Execute()
//
// KEY POINTS:
//   - Small, focused interfaces are easier to implement and compose
//   - Interface composition builds complex behaviors from simple parts
//   - Type assertions enable optional capability checking
//   - Plugin architecture benefits greatly from interfaces
//   - Compile-time verification ensures implementations are correct
//   - Interfaces enable the strategy pattern for swappable behaviors
//   - Real systems often check for optional interfaces at runtime
// ============================================================================
