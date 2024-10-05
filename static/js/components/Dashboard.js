const Dashboard = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">Welcome, {{ user?.name }}!</h2>
                <p class="mb-4">You have successfully logged in and completed the 2FA setup.</p>
                <button @click="logout" class="btn">Logout</button>
            </div>
        </div>
    `,
 data() {
  return {
   user: JSON.parse(localStorage.getItem('user')),
  };
 },
 async mounted() {
  await this.checkUser();
 },
 methods: {
  async logout() {
   try {
    await axios.post('/api/logout', { email: this.user.email });
    localStorage.removeItem('user');
    this.$router.push('/login');
   } catch (error) {
    alert('Logout failed: ' + error.response.data);
   }
  },
  async checkUser() {
   try {
    const user = JSON.parse(localStorage.getItem('user'));
    if (user) {
     return;
    } else {
     alert('Login to access dashboard');
     this.$router.push('/login');
    }
   } catch (error) {
    alert('Failed to get localstorage user');
   }
  },
 },
};
