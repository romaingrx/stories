new :
	touch main.html
	echo "<head>" >> main.html
	echo "    <title></title>" >> main.html
	echo "    <meta name="keywords" content="tag 1, tag 2"/>" >> main.html
	echo "</head>" >> main.html

publish :
	go run publisher/medium.go $(filter-out $@,$(MAKECMDGOALS))
