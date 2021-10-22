# Atividade 03

## Questão 01

### Item 0
A = B + C

```
<assign> => <id> = <expr>
=> A = <expr>
=> A = <expr> + <term>
=> A = <term> + <term>
=> A = <factor> + <term>
=> A = <id> + <term>
=> A = B + <term>
=> A = B + <factor>
=> A = B + <id>
=> A = B + C
```

![Árvore do Item 0](arvore.png)


## Questão 02

## Questão 03
```
for (int i = 0; i < 10; i++) {
   printf("%d\n", i);
   a = b + c;
}
<for_stmt> → for <control> { <stmt_list> }
<control> → (<init>; <expr_bool>; <expr_inc>)
<init> → <type> <var> = <value>
<stmt_list> → <stmt>; <stmt> | e
```




































