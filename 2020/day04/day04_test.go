package main

import (
	"testing"
)

func TestValidator(t *testing.T) {
	valid := validate("byr:2020")
	if valid {
		t.Fatalf("expected invalid")
	}

	valid = validate("byr:2002 iyr:2010 eyr:2020 hgt:193cm hcl:#aaaaaa ecl:amb pid:239239239")
	if !valid {
		t.Fatalf("expected valid")
	}

	valid = validate("byr:2003 iyr:2010 eyr:2020 hgt:193cm hcl:#aaaaaa ecl:amb pid:239239239")
	if valid {
		t.Fatalf("expected invalid")
	}

	valid = validate("byr:1900 iyr:2010 eyr:2020 hgt:193cm hcl:#aaaaaa ecl:amb pid:239239239")
	if valid {
		t.Fatalf("expected invalid")
	}

	valid = validate("byr:2000 iyr:2010 eyr:2020 hgt:58in hcl:#aaaaaa ecl:amb pid:239239239")
	if valid {
		t.Fatalf("expected invalid")
	}
}
