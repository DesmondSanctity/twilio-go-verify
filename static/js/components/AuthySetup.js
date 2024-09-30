const AuthySetup = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">Authy Setup</h2>
                <div v-if="!showTOTPChallenge">
                  <img v-if="qrCodeUrl" :src="qrCodeUrl" alt="QR Code" class="mb-4">
                  <p v-if="!qrCodeUrl" class="mb-4">Loading QR Code...</p>
                </div>
                <form @submit.prevent="verifyAuthy">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="code">Enter Authy Code</label>
                        <input class="form-input" id="code" v-model="code" type="text" required>
                    </div>
                    <button v-if="!user.TOTPEnabled" class="btn" type="submit">Enable TOTP</button>
                <button v-else class="btn" @click.prevent="createChallenge">Verify Challenge</button>
                </form>
            </div>
        </div>
    `,
 data() {
  return {
   qrCodeUrl: '',
   code: '',
   factorSid: '',
   user: JSON.parse(localStorage.getItem('user')),
   showTOTPChallenge: JSON.parse(localStorage.getItem('user')).TOTPEnabled,
  };
 },
 async mounted() {
  if (!this.user.TOTPEnabled) {
   await this.createTOTPFactor();
  }
 },
 methods: {
  async createTOTPFactor() {
   try {
    const user = JSON.parse(localStorage.getItem('user'));
    const response = await axios.post('/api/verify/create-totp', {
     email: user.email,
    });
    console.log(response.data);
    const qr = await qrcode(0, 'M');
    qr.addData(response.data.qrCode);
    qr.make();
    this.qrCodeUrl = qr.createDataURL(4);
    this.factorSid = response.data.factorSid;
   } catch (error) {
    alert('Failed to create TOTP factor: ' + error.response.data);
   }
  },
  async verifyAuthy() {
   try {
    const user = JSON.parse(localStorage.getItem('user'));
    await axios.post('/api/verify/verify-factor', {
     email: user.email,
     code: this.code,
    });
    this.user.TOTPEnabled = true;
    localStorage.setItem('user', JSON.stringify(this.user));
    this.$router.push('/dashboard');
   } catch (error) {
    alert('Authy verification failed: ' + error.response.data);
   }
  },
  async createTOTPChallenge() {
   try {
    await axios.post('/api/verify/create-totp-challenge', {
     email: this.user.email,
    });
    this.showTOTPChallenge = true;
   } catch (error) {
    alert('Failed to create TOTP challenge: ' + error.response.data);
   }
  },
 },
};
