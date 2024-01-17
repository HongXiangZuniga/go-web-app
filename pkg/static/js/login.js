const app = Vue.createApp({
  delimiters: ["{%", "%}"],
  data() {
    return {
      count: 0,
      Email: "",
      Password: "",
    };
  },
  methods: {
    async login() {
      if (!this.Email || !this.Password) {
        Swal.fire({
          icon: "error",
          title: "Error",
          text: "Please enter your email and password.",
        });
        return;
      }
      try {
        const response = await fetch("/auth", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            user: this.Email,
            password: this.Password,
          }),
        });
        if (response.ok) {
          window.location.href = "/profile";
        } else {
          const data = await response.json();
          Swal.fire({
            icon: "error",
            title: "Error",
            text: data.error || "Unexpected error",
          });
        }
      } catch (error) {
        console.error("Error making request:", error);
      }
    },
    handleEnterKey(event) {
      if (event.key === 'Enter') {
        this.login();
      }
    },
  },
  mounted() {
    // Add a listener for the keyup event on the document
    document.addEventListener('keyup', this.handleEnterKey);
  },
  beforeUnmount() {
    // Remove the listener when the Vue instance is destroyed
    document.removeEventListener('keyup', this.handleEnterKey);
  },
});

app.mount("#app");
