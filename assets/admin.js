function logout() {
	const cookies = document.cookie.split(";");

	for (let i = 0; i < cookies.length; i++) {
		const cookie = cookies[i];
		const eqPos = cookie.indexOf("=");
		const name = eqPos > -1 ? cookie.slice(0, eqPos) : cookie;
		document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
	}

	location.href = "/";
}

async function saveProfile() {
	const name = document.querySelector("#name").value;
	const email = document.querySelector("#email").value;
	const profession = document.querySelector("#profession").value;
	const imageUrl = document.querySelector("#imag_url").value;
	const linkedin = document.querySelector("#linkedin").value;
	const github = document.querySelector("#github").value;

	const response = await fetch("/api/user", {
		method: "PUT",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			name,
			email,
			profession,
			imageUrl,
			linkedin,
			github,
		}),
	});

	if (response.status !== 200) {
		alert("Failed to save profile");
	}

	location.reload();
}

function deletePost(id) {
	const modal = document.querySelector("#modal");
	modal.classList.remove("hidden");

	document
		.querySelector("#modal-cancel-button")
		.addEventListener("click", () => {
			modal.classList.add("hidden");
		});

	document
		.querySelector("#modal-delete-button")
		.addEventListener("click", async () => {
			const response = await fetch(`/api/blog/${id}`, {
				method: "DELETE",
			});

			if (response.status !== 204) {
				alert("Failed to delete post");
			}

			location.reload();
		});
}

function search(value) {
	urlParams = new URLSearchParams(window.location.search);
	if (value) urlParams.set("search", value);
	else urlParams.delete("search");
	location.search = urlParams;
}
