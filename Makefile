##
## EPITECH PROJECT, 2018
## 204ducks
## File description:
## Makefile for 204ducks.
##

SRC		=	main.go

SRC		:=	$(addprefix src/, $(SRC))

NAME	=	204ducks

all:	$(NAME)

$(NAME): $(SRC)
	go build -o $(NAME) $(SRC)

fclean:
	$(RM) $(NAME)

re: fclean all

.PHONY: all fclean re
