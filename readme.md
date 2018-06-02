# Wprowadzenie

## Język Go

- Stworzony i promowany przez Google
- Twórcy to m.in. Rob Pike, Robert Griesmier, Ken Thompson -
"ojcowie" UNIX'a, języków B i C, Java HotSpot, silnika V8 czy standardu UTF-8
- Na licencji open source
- Bardzo wydajny
    - Zaprojektowany tak aby wykorzystać możliwości współczesnych komputerów
    - Efektywne zarządzanie zależnościami
    - Szybki kompilator
- Stosunkowo prosty
- Nastawiony na pracę w zespołową (łatwe udostępnianie i wykorzystanie zewnętrznych pakietów)
- Stosowany przez wiele znanych firm np. Twitter, Apple, Amazon, Uber


## Wybrane cechy

- Statycznie typowany
- Kompilowany do kodu natywnego
- Stawia na uproszczony system typów i kompozycję zamiast dziedziczenia, oferuje własną
implementację/pomysł na interfejsy (duck typing)
- Oferuje uproszczony (z punktu widzenia programisty) i bezpieczny model tworzenia aplikacji wielowątkowych
- Automatyczne zarządzanie pamięcią (GC)
- Zredukowana składnia (mała ilość słów kluczowych)
- Kompletna biblioteka standardowa

## Zadanie praktyczne

- Instalacja i konfiguracja środowiska
- Korzystanie z dokumentacji
    - https://golang.org/doc
    - https://golang.org/pkg
    - https://golang.org/pkg/builtin
    - https://golang.org/src
- Użycie Go Playground (wykonywanie i współdzielenie kodu)
- Repozytorium pakietów https://godoc.org/    
- Stworzenie i uruchomienie pierwszej aplikacji    
    
## Podstawy tworzenia aplikacji

### Workspace

- Środowisko pracy
- Domyślna lokalizacja zależy od systemu operacyjnego (wartość zmienej GOPATH)
- Zawiera katalogi
    - **src** - pliki źródłowe pogrupowane w pakiety
    - **pkg** - zbudowane archiwa/biblioteki pogrupowane w pakiety
    - **bin** - zbudowane pliki wykonywalne

### Wybrane narzędzia dostępne z linii poleceń

- **go build** - budowanie pliku wykonywalnego/projektu np. go build github.com/landrzejewski/GoTraining/...
- **go run main.go** - zbudowanie i uruchomienie programu
- **go clean** - wyrzucenie pliku binarnego
- **go install** - przeniesienie pliku binarnego do katalogu bin
- **go get** - zaimportowanie zdalnego pakietu np. go get github.com/spf13/viper
- **go vet** - sprawdzenie kodu pod względem typowych błędów
- **go fmt** - sformatowanie kodu źródłowego

### Pakiety

- Umożliwiają organizację kodu i zapewniają separację przestrzeni nazewniczych
- Deklaracja przynależności do pakietu odbywa się na początku pliku źródłowego (nie licząc znaków białych i komentarzy)
- Nazwa katalogu powinna odpowiadać nazwie pakietu
- Zgodnie z dobrą praktyką pakiety powinny być zakładane wedłu wzoru repozytorium/nazwa_użytkownika/nazwa_pakietu
- Każda aplikacja musi zadeklarować metodę main w pakiecie main - w innym przypadku nie da się jej zbudować i uruchomić
- Nazwy pakietów powinny być krótkie i sensowne, piszemy je małymi literami
- Unikalność nazw pakietów nie jest wymagana - przy imporcie podajemy pełną ścieżkę
- Podczas importu można dokonać aliasowania (domyślnie obowiązuje ostatni człon ścieżki)
- Wszystkie elementy pliku pisane wielką literą są automatycznie eksportowane jako publiczne
- W ramach jednego pakietu funkcje i zmienne są współdzielone

### Importowanie

- Umożliwia wykorzystanie elementów z innych pakietów
- Deklaracja import pojawia się po package
- Można importować wiele pakietów naraz używając nawiasów okrągłych lub kilka razy słowa import
- Każdy import powinien być wykorzystany (inaczej wystąpi bąłd kompilacji)
- Pakiety wyszukiwane są względem katalogu/ów określonych przez zmienną GOPATH (workspace) i GOROOT
- Można stosować nazwane importy (aliasy) w celu odróżnienia tak samo nazwanych pakietów
- Czasem wymagane jest użycie pustego aliasu w postaci _ (wywołanie funkcji init bez jawnego użycia pakietu)

### Inicjalizacja pakietu

- Każdy pakiet może definiować dowolną ilość funkcji inicjalizacyjnych
uruchamianych przy starcie przed metodą main
- Mogą służyć do inicjalizacji lub wykonania jakichś zadań np. rejestracja
sterowników do bazy w pakiecie sql, który nic o nich nie wie
- Funkcje muszą nazywać się init i są wykonywane w kolejności importowania (tylko raz)

## Zadanie praktyczne

- Tworzenie i importowanie pakietów
- Publikowanie / ukrywanie elementów pakietu
- Przegląd najważniejszych elementów języka i konwencji
    - https://tour.golang.org
    - https://golang.org/doc/effective_go.html
- Analiza/omówienie przykładów
- Zadania do samodzielnego wykonania
