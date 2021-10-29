# Atividade 04
## João Marcelo

### Questão 01

```
<S> → <A>
<A> → <A> + <A> | <id>
<id> → a | b | c

a+a+b
<S> => <A>
=> <A> + <A>
=> <id> + <A>
=> a + <A>
=> a + <id>
=> a + b

```
### Questão 02
```
<program> → begin <stmt_list> end 
<stmt_list> → <stmt>
            | <stmt> ; <stmt_list> 
<stmt> → <var> = <expression>
<var> → A | B | C 

<expression> → <var> + <var>
<expression>.valor = <var>.valor + <var>.valor

             | <var> – <var> 
             | <var> * <var>
             | <var>
```

### Questão 03

### Questão 04

1. Regra sintâtica: ``` <expression> → <var> + <var> ```

   Regra semântica: ``` <expression> → <var> + <var> ```
2. 
3.
4.





