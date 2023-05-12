# Programarea in retea

Acest proiect reprezinta implementarea unei aplicatii de tip chat pentru obiectul **Programarea in Retea** in cadrul studiilor mele la **UTM**.

## Instalarea

Pentru a instala acest proiect este necesar de instalat pe calculator acest repository prin **git clone**, iar apoi de mers in acest directory si de instantiat
cele doua proiecte client si server. O data ce ne aflam in fiecare directorie, este necesar de rulat **go mod tidy** pentru a instala dependentele.

**Important**: Acest cod ruleaza cu **go1.20**

## Rularea

O data ce avem aplicatia instalata si gata de pornire, trebuie sa pornim ambele aplicatii in diferite terminaluri. Acest lucru se poate de facut prin comanda **go run .** in fiecare terminal care e deschis in ambele directorii.
Este important de mentionat ca ordinea de pornire este foarte importanta, astfel este necesar la inceput de pornit aplicatia server si apoi pe cea client. Mai multe instante de client sunt permise. Mai mult de atat ca fiecare
client sa aiba un nume de utilizator, este posibil de transmis acest nume ca argument la rularea aplicatiei prin **go run . numele-meu-de-utilizator**.

## Modul de lucru

Aplicatia server porneste un HTTP server care poate primi conexiuni dupa protocolul HTTP cat si WebSocket, intrucat noi suntem interesati in cel din urma, ne vom axa doar pe el. El accepta doar conexiuni websocket directe, fara un browser la mijloc,
pentru simplitate. O data ce o conexiune este primita, Handshake-ul se executa automat cu clientul. La fiecare conexiune, in memorie se salveaza client-ul ca ulterior sa putem face broadcast la mesajele venite.

Cat despre client, el instantiaza o conexiune la server utilizand protocolul WebSocket, prin **ws://** si transmite mesaje atunci cand un text este scris in terminal, precum si afiseaza orice text ce este transmis din partea server-ului.

## Demo

[Demo Video](./assets/demo.mov)
