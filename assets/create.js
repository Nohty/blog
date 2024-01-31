async function saveBlog() {
	const title = document.querySelector("#title").value;
	const category = document.querySelector("#category").value;
	const image = document.querySelector("#image").value;
	const published = document.querySelector("#published").checked;
	const content = document.querySelector("#content").value;

	const blogId = window.location.pathname.split("/")[3];

	if (blogId === undefined) {
		const response = await fetch("/api/blog", {
			method: "POST",
			body: JSON.stringify({ title, category, image, published, content }),
			headers: { "Content-Type": "application/json" },
		});

		if (response.status === 201) {
			return (location.href = "/admin");
		}
	} else {
		const response = await fetch(`/api/blog/${blogId}`, {
			method: "PUT",
			body: JSON.stringify({ title, category, image, published, content }),
			headers: { "Content-Type": "application/json" },
		});

		if (response.status === 200) {
			document.querySelector("#toast-success").classList.remove("hidden");

			document
				.querySelector("#toast-success button")
				.addEventListener("click", () => {
					document.querySelector("#toast-success").classList.add("hidden");
				});

			return setTimeout(() => {
				document.querySelector("#toast-success").classList.add("hidden");
			}, 3000);
		}
	}

	alert("Something went wrong");
}
