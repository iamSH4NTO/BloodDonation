<template>
  <footer class="bg-white pt-16 pb-8 border-t border-gray-100">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-12 mb-16">
        
        <!-- Brand -->
        <div class="space-y-4">
           <div class="flex items-center gap-2">
            <span class="material-icons text-red-600 text-2xl">bloodtype</span>
            <span class="text-xl font-bold text-gray-900 tracking-tight">BloodCare</span>
          </div>
          <p class="text-gray-500 text-sm leading-relaxed">
            BloodCare is an automated blood service that connects blood searchers with voluntary blood donors in a moment through SMS and location.
          </p>
          <div class="flex gap-4">
            <a href="#" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-500 hover:bg-red-50 hover:text-red-600 transition-colors">
                <span class="font-bold text-xs">f</span>
            </a>
            <a href="#" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-500 hover:bg-red-50 hover:text-red-600 transition-colors">
                <span class="font-bold text-xs">t</span>
            </a>
            <a href="#" class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-500 hover:bg-red-50 hover:text-red-600 transition-colors">
                <span class="font-bold text-xs">in</span>
            </a>
          </div>
        </div>

        <!-- Quick Links -->
        <div>
          <h3 class="font-bold text-gray-900 mb-6">Quick Links</h3>
          <ul class="space-y-3 text-sm text-gray-500">
            <li><a href="#" class="hover:text-red-600 transition-colors">About Us</a></li>
            <li><router-link to="/search" class="hover:text-red-600 transition-colors">Find a Donor</router-link></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">Privacy Policy</a></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">Terms of Service</a></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">Blog</a></li>
          </ul>
        </div>

        <!-- Support -->
        <div>
          <h3 class="font-bold text-gray-900 mb-6">Support</h3>
           <ul class="space-y-3 text-sm text-gray-500">
            <li><a href="#" class="hover:text-red-600 transition-colors">Contact Us</a></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">FAQs</a></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">Donate Funds</a></li>
            <li><a href="#" class="hover:text-red-600 transition-colors">Report an Issue</a></li>
          </ul>
        </div>

        <!-- Contact Us -->
        <div>
           <h3 class="font-bold text-gray-900 mb-6">Contact Us</h3>
           <ul class="space-y-4 text-sm text-gray-500">
             <li class="flex items-start gap-3">
               <span class="material-icons text-red-600 text-sm mt-0.5">location_on</span>
               <span>123 Medical Drive, Health City, HC 10001</span>
             </li>
             <li class="flex items-center gap-3">
               <span class="material-icons text-red-600 text-sm">email</span>
               <a href="mailto:help@BloodCare.com" class="hover:text-red-600">help@BloodCare.com</a>
             </li>
             <li class="flex items-center gap-3">
               <span class="material-icons text-red-600 text-sm">call</span>
               <a href="tel:+8801234567890" class="hover:text-red-600">+880 1234 567 890</a>
             </li>
           </ul>
        </div>

      </div>

      <div class="border-t border-gray-100 pt-8 flex flex-col md:flex-row justify-between items-center gap-4 text-xs text-gray-400">
        <p>&copy; 2023 BloodCare. All rights reserved.</p>
        <p class="flex items-center gap-1">Made with <span class="material-icons text-red-600 text-xs">favorite</span> for humanity.</p>
      </div>
    </div>
  </footer>
</template>
