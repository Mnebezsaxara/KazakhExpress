document.addEventListener("DOMContentLoaded", () => {
  const AUTH_API_URL = "http://localhost:8082"; // Auth service URL
  const loginTab = document.getElementById("login-tab");
  const signupTab = document.getElementById("signup-tab");
  const loginForm = document.getElementById("login-form");
  const signupForm = document.getElementById("signup-form");

  loginTab.addEventListener("click", () => {
    loginTab.classList.add("active");
    signupTab.classList.remove("active");
    loginForm.classList.add("active");
    signupForm.classList.remove("active");
  });

  signupTab.addEventListener("click", () => {
    signupTab.classList.add("active");
    loginTab.classList.remove("active");
    signupForm.classList.add("active");
    loginForm.classList.remove("active");
  });

  // Handle Login
  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;
    const message = document.getElementById("login-message");

    try {
      const res = await fetch(`${AUTH_API_URL}/api/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      const data = await res.json();

      if (res.ok) {
        // Store the token in localStorage
        if (data.token) {
          localStorage.setItem("token", data.token);
          message.textContent = "Login successful!";
          message.classList.remove("auth-error");
          // Show success alert and redirect
          alert("Login successful! Welcome back!");
          window.location.href = "/public/index.html";
        }
      } else {
        message.textContent = data.message || "Login failed.";
        message.classList.add("auth-error");
      }
    } catch (error) {
      message.textContent = "Error connecting to the server.";
      message.classList.add("auth-error");
      console.error("Login error:", error);
    }
  });

  // Handle Sign Up
  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("signup-name").value;
    const email = document.getElementById("signup-email").value;
    const password = document.getElementById("signup-password").value;
    const message = document.getElementById("signup-message");

    try {
      const res = await fetch(`${AUTH_API_URL}/api/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, password }),
      });

      const data = await res.json();

      if (res.ok) {
        message.textContent = "Signup successful! Please login.";
        message.classList.remove("auth-error");
        // Show success alert
        alert("Sign up successful! Please log in with your credentials.");
        // Switch to login tab after successful signup
        loginTab.click();
        // Clear signup form
        signupForm.reset();
      } else {
        message.textContent = data.message || "Signup failed.";
        message.classList.add("auth-error");
      }
    } catch (error) {
      message.textContent = "Error connecting to the server.";
      message.classList.add("auth-error");
      console.error("Signup error:", error);
    }
  });
});
