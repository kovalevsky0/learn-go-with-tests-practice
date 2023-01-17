package dictionary

import "testing"

func assertStrings(t *testing.T, result, expected string) {
	t.Helper()

	if expected != result {
		t.Errorf("Expected %q but received %q", expected, result)
	}
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"agent": "John Smith"}

	t.Run("existing word", func(t *testing.T) {
		result, _ := dictionary.Search("agent")
		expected := "John Smith"

		assertStrings(t, result, expected)
	})

	t.Run("not existing word", func(t *testing.T) {
		_, err := dictionary.Search("something")

		if err == nil {
			t.Fatal("Expected an error but received nil")
		}

		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"fireball": "fire magic"}

	t.Run("new word", func(t *testing.T) {
		dictionary.Add("stupefy", "stunning spell")
		result, _ := dictionary.Search("stupefy")
		expected := "stunning spell"

		assertStrings(t, result, expected)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add("fireball", "something")

		assertStrings(t, err.Error(), ErrWordExists.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "stupefy"
		definition := "stunning spell"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		err := dictionary.Update(word, definition)
		result, _ := dictionary.Search(word)

		if err != nil {
			t.Fatal("Expected nil but received an error")
		}
		assertStrings(t, result, definition)
	})
	t.Run("not existing word", func(t *testing.T) {
		word := "stupefy"
		definition := "stunning spell"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)

		if err == nil {
			t.Fatal("Expected an error but received nil")
		}

		assertStrings(t, err.Error(), ErrWordDoesNotExist.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "stupefy"
		definition := "stunning spell"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		err := dictionary.Delete(word)
		_, errSearch := dictionary.Search(word)

		if err != nil {
			t.Fatal("Expected nil but received an error")
		}

		assertStrings(t, errSearch.Error(), ErrNotFound.Error())
	})
	t.Run("not existing word", func(t *testing.T) {
		word := "stupefy"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		if err == nil {
			t.Fatal("Expected an error but received nil")
		}

		assertStrings(t, err.Error(), ErrWordDoesNotExist.Error())
	})
}
