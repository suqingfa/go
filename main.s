TEXT ·add(SB), $8-24
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    ADDQ BX, AX
    MOVQ AX, (SP)
    MOVQ (SP), AX
    MOVQ AX, ret+16(FP)
    RET
