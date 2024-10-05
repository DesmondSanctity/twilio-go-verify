const Signup = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">Sign Up</h2>
                <form @submit.prevent="signup">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="name">Name</label>
                        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        id="name" v-model="name" type="text" required>
                    </div>
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
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="phone">Phone Number</label>
                        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        id="phone" v-model="phone" type="tel" required>
                    </div>
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                    type="submit">Sign Up</button>
                </form>
            </div>
        </div>
    `,
 data() {
  return {
   name: '',
   email: '',
   password: '',
   phone: '',
  };
 },
 methods: {
  async signup() {
   try {
    await axios.post('/api/signup', {
     name: this.name,
     email: this.email,
     password: this.password,
     phoneNumber: this.phone,
    });
    alert('Signup successful! Please login.');
    this.$router.push('/login');
   } catch (error) {
    alert('Signup failed: ' + error.response.data);
   }
  },
 },
};
