const Login = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">Login</h2>
                <form @submit.prevent="login">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="email">Email</label>
                        <input class="form-input" id="email" v-model="email" type="email" required>
                    </div>
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="password">Password</label>
                        <input class="form-input" id="password" v-model="password" type="password" required>
                    </div>
                    <button class="btn" type="submit">Login</button>
                </form>
            </div>
        </div>
    `,
 data() {
  return {
   email: '',
   password: '',
  };
 },
 methods: {
  async login() {
   try {
    const response = await axios.post('/api/login', {
     email: this.email,
     password: this.password,
    });
    localStorage.setItem('user', JSON.stringify(response.data));
    await this.sendSMSOTP();
    this.$router.push('/sms-verification');
   } catch (error) {
    alert('Login failed: ' + error);
   }
  },
  async sendSMSOTP() {
   try {
    const user = JSON.parse(localStorage.getItem('user'));
    await axios.post('/api/verify/send-sms', {
     email: user.email,
    });
   } catch (error) {
    alert('Failed to send SMS OTP: ' + error);
   }
  },
 },
};
