<script setup lang="ts">
  import { RouterLink } from 'vue-router'
  // defineProps<{
  // }>()
  </script>

<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark" aria-label="Ninth navbar example">
    <div class="container-xl">
      <a class="navbar-brand" href="/">一介白衣</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarsExample07XL" aria-controls="navbarsExample07XL" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExample07XL">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <RouterLink to="/article" class="nav-link" href="#">博文</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink to="/poem" class="nav-link" href="#">小诗</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink to="/essay" class="nav-link" href="#">随笔</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink to="/novelette" class="nav-link" href="#">轻小说</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink to="/about" class="nav-link" href="#">关于</RouterLink>
          </li>
        </ul>
        <form role="search">
          <input class="form-control" type="search" placeholder="Search" aria-label="Search">
        </form>
      </div>
    </div>
  </nav>

  <!-- <div class="container-xl mb-4">
    <p>Matching .container-xl...</p>
  </div> -->


  <!-- <div class="container">
    <div>
      <div class="bg-light p-5 rounded">
        <div class="col-sm-8 mx-auto">
          <h1>Navbar examples</h1>
          <p>This example is a quick exercise to illustrate how the navbar and its contents work. Some navbars extend the width of the viewport, others are confined within a <code>.container</code>. For positioning of navbars, checkout the <a href="../examples/navbar-static/">top</a> and <a href="../examples/navbar-fixed/">fixed top</a> examples.</p>
          <p>At the smallest breakpoint, the collapse plugin is used to hide the links and show a menu button to toggle the collapsed content.</p>
          <p>
            <a class="btn btn-primary" href="../components/navbar/" role="button">View navbar docs &raquo;</a>
          </p>
        </div>
      </div>
    </div>
  </div> -->
</template>
<style>
  .bd-placeholder-img {
    font-size: 1.125rem;
    text-anchor: middle;
    -webkit-user-select: none;
    -moz-user-select: none;
    user-select: none;
  }

  @media (min-width: 768px) {
    .bd-placeholder-img-lg {
      font-size: 3.5rem;
    }
  }

  .b-example-divider {
    height: 3rem;
    background-color: rgba(0, 0, 0, .1);
    border: solid rgba(0, 0, 0, .15);
    border-width: 1px 0;
    box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
  }

  .b-example-vr {
    flex-shrink: 0;
    width: 1.5rem;
    height: 100vh;
  }

  .bi {
    vertical-align: -.125em;
    fill: currentColor;
  }

  .nav-scroller {
    position: relative;
    z-index: 2;
    height: 2.75rem;
    overflow-y: hidden;
  }

  .nav-scroller .nav {
    display: flex;
    flex-wrap: nowrap;
    padding-bottom: 1rem;
    margin-top: -1px;
    overflow-x: auto;
    text-align: center;
    white-space: nowrap;
    -webkit-overflow-scrolling: touch;
  }
</style>
<!-- <link href="../assets/dist/css/bootstrap.min.css" rel="stylesheet" /> -->
<!-- <link rel="canonical" href="https://getbootstrap.com/docs/5.2/examples/navbars/" /> -->
<!-- <link href="navbar.css" rel="stylesheet" /> -->
<!-- <script src="../assets/dist/js/bootstrap.bundle.min.js"></script> -->


