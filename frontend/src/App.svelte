<script lang="ts">
  import Feed from "./pages/Feed.svelte";
  import { Route, Router } from "svelte-navigator";
  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
  import Header from "./lib/Header.svelte";
  import User from "./pages/User.svelte";
</script>

<Router>
  <Header />
  <Route path="/" component={Feed} />
  <Route path="/login" component={Login} />
  <Route path="/register" component={Register} />
  <Route path="/user/:id" let:params>
    <User id={params.id} />
  </Route>
</Router>
