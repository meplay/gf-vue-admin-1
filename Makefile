develop-merge-master:
	@git checkout develop
	@git pull origin master
	@git add .
	@git push origin develop
	@git checkout master

gfva_develop-merge-master:
	@git checkout gfva_develop
	@git pull origin master
	@git add .
	@git push origin gfva_develop
	@git checkout master

help:
	@echo "make develop-merge-master - develop从master拉去代码并合并, 合并完成返回master分支"
	@echo "make gfva_develop-merge-master - gfva_develop从master拉去代码并合并, 合并完成返回master分支"