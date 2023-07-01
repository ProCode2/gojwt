<script lang="ts">
  import Counter from './lib/Counter.svelte'
</script>

<main class="w-screen h-screen">
  <div class="w-full h-full">
    <div class="w-full h-full overflow-auto no-scrollbar snap-y snap-mandatory scroll-smooth">
      <div class="snap-always snap-center h-full p-2 flex justify-center items-center flex-col text-center">
        <span class="text-2xl font-bold tracking-wider">We are changing how text content is consumed</span>
        <span>Once and for All.</span>
      </div>
      <div class="snap-always snap-center h-full p-2">And its not what you think!</div>
      <div class="snap-always snap-center h-full p-2">
        So you might want you join the ride.
      </div>
    </div>
  </div>
</main>

<style>
</style>
