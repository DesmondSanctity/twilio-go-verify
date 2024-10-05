const Login = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">Login</h2>
                <form @submit.prevent="login">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="email">Email</label>
                        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        id="email" v-model="email" type="email" required>
                    </div>
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="password">Password</label>
                        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        id="password" v-model="password" type="password" required>
                    </div>
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                    type="submit">Login</button>
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
