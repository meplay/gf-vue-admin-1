ll: master-merge-from-gfva_develop develop-merge-from-master gfva_develop

master:
	@git checkout master

develop:
	@git checkout develop

gfva_develop:
	@git checkout gfva_develop

develop-merge-from-master:
	@git checkout develop
	@git pull origin master
	@git add .
	@git push origin develop
	@git checkout master

gfva_develop-merge-from-master:
	@git checkout gfva_develop
	@git pull origin master
	@git add .
	@git push origin gfva_develop
	@git checkout master

master-merge-from-develop:
	@git checkout master
	@git pull origin develop
	@git add .
	@git push origin master

master-merge-from-gfva_develop:
	@git checkout master
	@git pull origin gfva_develop
	@git add .
	@git push origin master

help:
	@echo "make - gfva_develop代码合并到master和develop, 合并完成返回gfva_develop分支!"
	@echo "make develop-merge-from-master - develop从master拉去代码并合并, 合并完成返回master分支!"
	@echo "make gfva_develop-merge-from-master - gfva_develop从master拉去代码并合并, 合并完成返回master分支!"
	@echo "make master-merge-from-develop - master从develop拉去代码并合并!"
	@echo "make master-merge-from-gfva_develop - master从gfva_develop拉去代码并合并!"