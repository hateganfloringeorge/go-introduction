package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new entry", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "new"
		definition := "testing new entry"

		err := dictionary.Add("new", "testing new entry")

		if err != nil {
			t.Fatal("got error when not expected")
		}
		assertDictionary(t, dictionary, word, definition)
	})

	t.Run("add already existing word", func(t *testing.T) {
		word := "old"
		oldDefinition := "old entry"
		dictionary := Dictionary{word: oldDefinition}

		err := dictionary.Add(word, "new entry")

		assertError(t, err, ErrWordExists)
		assertDictionary(t, dictionary, word, oldDefinition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update entry", func(t *testing.T) {
		word := "key"
		oldDefinition := "old definition"
		newDefinition := "new definition"
		dictionary := Dictionary{word: oldDefinition}

		err := dictionary.Update(word, newDefinition)
		
		assertNoError(t, err)
		assertDictionary(t, dictionary, word, newDefinition)
	})

	t.Run("no entry to update", func(t *testing.T) {
		word := "key"
		definition := "entry value"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrUpdateFailed)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete entry", func(t *testing.T) {
		word := "key"
		definition := "definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})
}

// func assertNoEntry(t *testing.T, dictionary Dictionary, word string) {

	
// }

func assertDictionary(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s wanted %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got error when not expected")
	}
}