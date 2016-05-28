package spec

import (
	"testing"

	"github.com/alligrader/gradebook-backend/tasks"
)

func TestParseXML(t *testing.T) {

	var doc *tasks.CheckstylePayload = tasks.ParseXMLDocument(checkstyle_output)

	if doc.XMLName.Local != "checkstyle" {
		t.Fatal("XMLName does not match expected string checkstyle")
	}

	if doc.Version != "6.17" {
		t.Fatal("Version does not match expected string 6.17")
	}

	if doc.File.XMLName.Local != "file" {
		t.Fatal("File XMLName does not match expected string file")
	}

	if doc.File.Name != "/usr/src/myapp/Main.java" {
		t.Fatal("File Name does not match expected string /usr/src/myapp/Main.java")
	}

	if len(doc.File.Errors) < 5 {
		t.Fatal("Number of errors detected less than expected number of 5")
	}

	if len(doc.File.Errors) > 5 {
		t.Fatal("Number of errors detected greater than expected number of 5")
	}

	if doc.File.Errors[0].Line != "5" {
		t.Fatal("File error line does not match expected value of 5")
	}

	if doc.File.Errors[1].Line != "5" {
		t.Fatal("File error line does not match expected value of 5")
	}

	if doc.File.Errors[2].Line != "24" {
		t.Fatal("File error line does not match expected value of 5")
	}

	if doc.File.Errors[3].Line != "39" {
		t.Fatal("File error line does not match expected value of 39")
	}

	if doc.File.Errors[4].Line != "50" {
		t.Fatal("File error line does not match expected value of 50")
	}

	if doc.File.Errors[1].Column != "5" {
		t.Fatal("File error column does not match expected value of 5")
	}

	if doc.File.Errors[0].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[1].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[2].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[3].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[4].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[0].Message != "'method def modifier' have incorrect indentation level 4, expected level should be 2." {
		t.Fatal("File Fatal Message does not match expected string 'method def modifier' have incorrect indentation level 4, expected level should be 2.")
	}

	if doc.File.Errors[1].Message != "Missing a Javadoc comment." {
		t.Fatal("File Fatal Message does not match expected string 'method def modifier' have incorrect indentation level 4, expected level should be 2.")
	}

	if doc.File.Errors[2].Message != "'for rcurly' have incorrect indentation level 8, expected level should be 4." {
		t.Fatal("File Fatal Message does not match expected string 'for rcurly' have incorrect indentation level 8, expected level should be 4.")
	}

	if doc.File.Errors[3].Message != "'if rcurly' have incorrect indentation level 12, expected level should be 6." {
		t.Fatal("File Fatal Message does not match expected string 'if rcurly' have incorrect indentation level 12, expected level should be 6.")
	}

	if doc.File.Errors[4].Message != "'method def rcurly' have incorrect indentation level 4, expected level should be 2." {
		t.Fatal("File Fatal Message does not match expected string 'method def modifier' have incorrect indentation level 4, expected level should be 2.")
	}
}

var checkstyle_output []byte = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<checkstyle version="6.17">
<file name="/usr/src/myapp/Main.java">
<error line="5" severity="warning" message="&apos;method def modifier&apos; have incorrect indentation level 4, expected level should be 2." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
<error line="5" column="5" severity="warning" message="Missing a Javadoc comment." source="com.puppycrawl.tools.checkstyle.checks.javadoc.JavadocMethodCheck"/>
<error line="24" severity="warning" message="&apos;for rcurly&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
<error line="39" severity="warning" message="&apos;if rcurly&apos; have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
<error line="50" severity="warning" message="&apos;method def rcurly&apos; have incorrect indentation level 4, expected level should be 2." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
</file>
</checkstyle>`)
