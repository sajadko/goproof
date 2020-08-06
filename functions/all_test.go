package functions

import "testing"

func TestIsEmail(t *testing.T) {
	result := IsEmail("abcd@abcd.abcd")
	if result == false {
		t.Errorf("IsEmail Failed. got: %v , want: %v", result, true)
	}
}

func TestMaxLength(t *testing.T) {
	result := MaxLength("abcdefg", "10")
	if result == false {
		t.Errorf("MaxLength Failed. got: %v , want: %v", result, true)
	}
}

func TestIsJSON(t *testing.T) {
	result := IsJSON(`{"glossary":{"title":"exampleglossary","GlossDiv":{"GlossList":{"GlossEntry":{"GlossTerm":"StandardGeneralizedMarkupLanguage","GlossDef":{"GlossSeeAlso":["GML","XML"]}}}}}}`)
	if result == false {
		t.Errorf("IsJSON Failed. got: %v , want: %v", result, true)
	}
}	

func TestHasWhiteSpace(t *testing.T) {
	result := HasWhiteSpace(`abcd e`)
	if result == false {
		t.Errorf("HasWhiteSpace Failed. got: %v , want: %v", result, true)
	}
}

func TestIsCreditCard(t *testing.T) {
	result := IsCreditCard(`2720993186358710`)
	if result == false {
		t.Errorf("IsCreditCard Failed. got: %v , want: %v", result, true)
	}
}

func TestIsDataURI(t *testing.T) {
	result := IsDataURI(`data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==`)
	if result == false {
		t.Errorf("IsDataURI Failed. got: %v , want: %v", result, true)
	}
}

func TestMultiple(t *testing.T) {
	result := Multiple(12,2)
	if result == false {
		t.Errorf("Multiple Failed. got: %v , want: %v", result, true)
	}
}