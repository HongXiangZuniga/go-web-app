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
          title: "Oops...",
          text: "Por favor, ingresa tu correo electrónico y contraseña.",
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
            title: "Oops...",
            text: data.error || "Error inesperado",
          });
        }
      } catch (error) {
        console.error("Error al realizar la solicitud:", error);
      }
    },
  },
});

app.mount("#app");
