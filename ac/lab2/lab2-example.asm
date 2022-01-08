; z=(a-b*c/d)/(c+2-a/b)+5
; z=(5*a-b/7)/(3/b + a * a)

INCLUDE Irvine32.inc
.data
a dw 5; cu o valoare de 5 cuvinte de tip pentru a
b db 1; cu o valoare de 1 octet un cuvânt pentru b
cd db 10; cu o valoare de 10 octeți un cuvânt pentru cd
d dw 5; cu o valoare de 5 cuvinte de tip pentru d
interm dd ?
rez dd ?
.code

main proc

mov eax, 0; punem valoarea 0 în registrul eax
mov al, b; punem valoarea b în registrul al
imul cd; rezultatul b* c este salvat în ax
cwd; extindem al la ax(ținând cont de semn)
idiv d; coeficientul în ax și restul în dx, ax = b * c / d
mov interm, eax; interm = b * c / d
sub a, ax; ax = a –b * c / d
cwd; extindem cuvântul din ax, într -un cuvânt dublu în dx : ax
mov interm, eax; interm = a –b * c / d
mov ax, a; scriem in ax
idiv b; executam a / b
cwd; extindem rezultatul împărțirii de la ax în eax(păstrând semnul)
mov ecx, eax; copiem rezultatul a / b în ecx
mov al, cd
cbw
add ax, 2
mov ebx, interm; copiem valoarea interm în registrul ebx
mov word ptr interm, ax; copiem in interm mic și mare cuvintele rezultate
mov word ptr interm[2], dx
sub interm, ecx; executam a / b -c + a
mov eax, ebx; copiem rezultatul a -b * c / d în registrul eax
cdq; extindem eax la edx : eax(luand in seama semnul)
idiv interm; executam(a -b * c / d) / (c + 2 -a / b)
mov ax, 5; scriem 5 in ax
add interm, eax; executam eax = (a -b * c / d) / (c + 2 -a / b) + 5
cbw; extindem rezultatul diviziunii de la al la ax(păstrând semnul)
cwd; extindem rezultatul împărțirii de la ax pana la eax(păstrând semnul)
mov rez, eax; copiem rezultatul în rez
Exit; Procedura de iesire din Irvine32
main ENDP; Terminarea procedurii main
END main; Punctul de terminare a programului