Vue.prototype.$axios = axios;

var header = new Vue({
    el: '#header',
    data: function() {
        return {
          visible: false,
          activeIndex: '1',
          activeIndex2: '1'
        }
    },
    methods: {
        handleSelect(key, url) {
          switch(url.length){
            case 1:
              location.href = url[0];
              break;
            case 2:
              location.href = url[1];
              break;
          }
        },
    }
})

var menu = new Vue({
  el: '#aside',
  data: function() {
      return {
      }
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
})
