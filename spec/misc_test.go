package spec

import (
	"testing"

	"github.com/alligrader/gradebook-backend/tasks"
)

func TestParseXML(t *testing.T) {
	s := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<checkstyle version="6.17">
	<file name="/usr/src/myapp/Main.java">
	<error line="5" severity="warning" message="&apos;method def modifier&apos; have incorrect indentation level 4, expected level should be 2." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="5" column="5" severity="warning" message="Missing a Javadoc comment." source="com.puppycrawl.tools.checkstyle.checks.javadoc.JavadocMethodCheck"/>
	<error line="7" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="9" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="11" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="13" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="15" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="16" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="17" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="19" severity="warning" message="&apos;for&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="19" column="12" severity="warning" message="WhitespaceAround: &apos;for&apos; is not followed by whitespace. Empty blocks may only be represented as {} when not part of a multi-block statement (4.1.3)" source="com.puppycrawl.tools.checkstyle.checks.whitespace.WhitespaceAroundCheck"/>
	<error line="20" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="21" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="22" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="23" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="24" severity="warning" message="&apos;for rcurly&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="26" severity="warning" message="&apos;for&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="26" column="12" severity="warning" message="WhitespaceAround: &apos;for&apos; is not followed by whitespace. Empty blocks may only be represented as {} when not part of a multi-block statement (4.1.3)" source="com.puppycrawl.tools.checkstyle.checks.whitespace.WhitespaceAroundCheck"/>
	<error line="27" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="28" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="29" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="30" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="31" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="32" severity="warning" message="&apos;for&apos; child have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="33" severity="warning" message="&apos;for rcurly&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="35" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="36" severity="warning" message="&apos;for&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="36" column="12" severity="warning" message="WhitespaceAround: &apos;for&apos; is not followed by whitespace. Empty blocks may only be represented as {} when not part of a multi-block statement (4.1.3)" source="com.puppycrawl.tools.checkstyle.checks.whitespace.WhitespaceAroundCheck"/>
	<error line="37" severity="warning" message="&apos;if&apos; have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="37" column="15" severity="warning" message="WhitespaceAround: &apos;if&apos; is not followed by whitespace. Empty blocks may only be represented as {} when not part of a multi-block statement (4.1.3)" source="com.puppycrawl.tools.checkstyle.checks.whitespace.WhitespaceAroundCheck"/>
	<error line="38" severity="warning" message="&apos;if&apos; child have incorrect indentation level 16, expected level should be 8." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="39" severity="warning" message="&apos;if rcurly&apos; have incorrect indentation level 12, expected level should be 6." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="40" severity="warning" message="&apos;for rcurly&apos; have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="42" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="44" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="45" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="46" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="47" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="48" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="49" severity="warning" message="&apos;method def&apos; child have incorrect indentation level 8, expected level should be 4." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	<error line="50" severity="warning" message="&apos;method def rcurly&apos; have incorrect indentation level 4, expected level should be 2." source="com.puppycrawl.tools.checkstyle.checks.indentation.IndentationCheck"/>
	</file>
	</checkstyle>`)

	var doc *tasks.CheckstylePayload = tasks.ParseXMLDocument(s)

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

	if len(doc.File.Errors) < 10 {
		t.Fatal("Number of errors detected less than 10")
	}

	if doc.File.Errors[0].Line != "5" {
		t.Fatal("File error line does not match expected value of 5")
	}

	if doc.File.Errors[0].Severity != "warning" {
		t.Fatal("Severity does not match expected string warning")
	}

	if doc.File.Errors[0].Message != "'method def modifier' have incorrect indentation level 4, expected level should be 2." {
		t.Fatal("File Fatals Message does not match expected string 'method def modifier' have incorrect indentation level 4, expected level should be 2.")
	}
}
