; .model small
; .stack 100h
.data
x DB 32h
y DB 08h
.code
Start:
mov ax, @data
mov ds,ax

mov al,y
mov bl,2
mul bl ; al = 2y
cmp x,al; comparam prin x - 2y > 0
JNGE DOIF1; Jump if not greater or equal

mov al,x
mov bl,4
mul bl ; al = 4x
sub al,y ; 4x - y
JMP L1 ; sari la L1
DOIF1: 
mov bh,y
mov bl,x
sub bh,bl ; y - x 

mov ax,0
mov al ,bh
mov bh,2h
div bh ; executam (y-x)/2
mov bh,102h ; executam (y-x)/2 + 102
add al,bh
L1:


end Start