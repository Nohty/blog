document.addEventListener("DOMContentLoaded", function () {
	document
		.querySelector("#login")
		.addEventListener("submit", async function (e) {
			e.preventDefault();

			const email = e.target.querySelector("#email").value;
			const password = e.target.querySelector("#password").value;

			const response = await fetch("/api/auth/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email, password }),
			});

			if (response.status !== 200) {
				return location.reload();
			}

			location.href = "/admin";
		});
});
