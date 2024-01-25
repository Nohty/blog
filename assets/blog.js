function filterByCategory(category) {
	urlParams = new URLSearchParams(window.location.search);
	if (category !== "All Categories") urlParams.set("category", category);
	else urlParams.delete("category");
	location.search = urlParams;
}

function search(value) {
	urlParams = new URLSearchParams(window.location.search);
	if (value) urlParams.set("search", value);
	else urlParams.delete("search");
	location.search = urlParams;
}
