# minyr

Minyr av Kjetil Nymoen

I main.go filen ligger innholdet til programmet. 
Her blir også yr.go funksonene ConvertTemperature 
og AverageTemperature kalt.

I yr.go finner vi funksjonsdefinisjoner for ConvertTemperature, AverageTemperature, og andre underfunksjoner. 
Den inkluderer funksjoner for å sjekke om en output-fil allerede eksisterer og for å opprette en ny output-fil hvis den ikke finnes. Den har også en funksjon for å beregne gjennomsnittstemperaturen i input-filen og konvertere den til Celsius eller Fahrenheit basert på brukerens valg.

I yr_test.go finner vi tester som teller antall linjer i filen, konverterer enkelte linjer til et ønsket format, og som beregner gjennomsnittstemperaturen for en bestemt tidsperiode i filen. Testene inkluderer også eksempler på ønsket input og forventet output.