; z=(5*a-b/7)/(3/b + a * a)

INCLUDE Irvine32.inc
.data
a db 0; initiem variabila a cu 3
b dw 1; initiem variabila a cu 3

cd db 5; initiem constanta c cu 5
d dw 7; initiem constanta c cu 7
e dw 3; initiem constanta c cu 3

interm dd ?
interm2 dd ?
rez dd ?
.code

main proc

mov eax, 0; punem valoarea 0 în registrul eax

mov al, a; punem valoarea a în registrul al
imul cd; al = 5 * a sau c * a
cbw; din al in ax
cwd; din ax in eax
mov interm, eax

mov ax, b
idiv d; ax = ax / d sau ax = b / 7
cwd; din ax in eax

sub interm, eax; (5*a-b/7)

mov ax, e
idiv b; 3 / b
cwd; din ax in eax
mov interm2, eax


mov al, a
imul a; a * a
cbw; din al in ax
cwd; din ax in eax

add interm2, eax; (3/b + a * a)

mov ebx, interm2
mov eax, interm
idiv ebx; executam impartirea intre 2 paranteze, astfel eax va contine rezultatul final
mov rez, eax


Exit; Procedura de iesire din Irvine32
main ENDP; Terminarea procedurii main
END main; Punctul de terminare a programului