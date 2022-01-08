%include Irvine32.inc
.data
Promt DB 'Doresti sa devii programator?[y/n]', 0;
Dad DB 13, 10, 'Vei deveni!', 13, 10, 0
Nud DB 13, 10, 'Vei deveni filosof!', 13, 10, 0
.code
main PROC; se indica procedura cu numele main
mov edx, OFFSET Promt; edx - deplasamentul(offset) sirului Promt
call WriteString; apeleaza functia WriteString
call ReadChar; apeleaza functia ReachChar
cmp al, 'y'; comparearea din al se scade codul ASCII al literei y
jz IsDad; salt conditionat(jz - jump if zero), daca rezultatul compararii este zero, salt la eticheta IsDad
cmp al, 'n'; comparearea din al se scade codul ASCII al literei n
jz IsNud; da rezultatul compararii este zero, salt la eticheta IsNud
IsDad : mov edx, OFFSET Dad; in edx - offsetul sirului Dad
call WriteString; apeleaza functia WriteString
jmp ex
IsNud : mov edx, OFFSET Nud; in edx - offsetul sirului Nud
call WriteString; apeleaza functia WriteString
ex : exit
main ENDP
END main