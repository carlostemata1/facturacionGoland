<template>
  <v-app>
    <!--Iconos gracias a: https://next.vuetifyjs.com/en/features/icon-fonts/#material-design-icons-->
    <v-navigation-drawer app id="navigation" v-model="drawer">
      <h2>Prueba tecnica</h2>
      <v-divider></v-divider>
      <v-list>
        <v-item-group>
          <div>
            <v-item>
              <router-link to="/"
                ><v-icon icon="mdi-home-circle-outline" /> Factura</router-link
              >
            </v-item>
          </div>
          <v-divider></v-divider>
          <div>
            <v-item>
              <router-link to="/promocion">
                <v-icon icon="mdi-package-variant-closed" />
                Promocion</router-link
              >
            </v-item>
          </div>
          <v-divider></v-divider>
          <div>
            <v-item>
              <router-link to="/factura"
                ><v-icon icon="mdi-clipboard-multiple-outline" />
                Pedidos</router-link
              >
            </v-item>
          </div>
          <v-divider></v-divider>
          <div>
            <v-item>
              <router-link to="/contacto"
                ><v-icon icon="mdi-cellphone-message" /> Contacto</router-link
              >
            </v-item>
          </div>
        </v-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app>
      <h1 @click.stop="drawer = !drawer">{{ this.$route.name }}</h1>
    </v-app-bar>

    <v-main @click.stop="drawer = null">
      <v-container fluid >
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: null,
    };
  },
};
</script>

<style lang="scss">
#navigation {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;

  div {
    margin: 15px 0;
  }

  a {
    color: #2c3e50;
    text-decoration: none;

    &:hover {
      color: #7bc6ff;
    }
  }
}
</style>
