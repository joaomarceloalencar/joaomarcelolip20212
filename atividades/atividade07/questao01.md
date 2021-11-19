# Atividade 07

Pilha     | Entrada         | Ação
--------- | --------------- | ------ 
0         | id * (id + id)$ | Shift 5
0id5      |   * (id + id) $ | Reduce 6
0F3       |   * (id + id) $ | Reduce 4
0T2       |   * (id + id) $ | Shift 7
0T2*7     |     (id + id) $ |

4. T -> F
Lado Direito (Beta): F
Lado Esquedo (A): T

6. F -> id
Lado Direito (Beta): id
Lado Esquerdo(A) : F
