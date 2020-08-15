<template>
    <div>
        <div style="width: 100%">
            <APlayer :audio="audio"></APlayer>
        </div>
    </div>
</template>

<script>
    import { APlayer } from '@moefe/vue-aplayer';
    export default {
        name: "musicPlayer",
        components:{
            APlayer
        },
        data(){
            return {
                audio :[
                    {
                        name: '东西（Cover：林俊呈）',
                        artist: '纳豆',
                        url: 'http://m10.music.126.net/20200815150040/7c6379465e0504ed837a11da3f2e89d4/ymusic/1643/12cd/ba47/2b900f22b1d5d7047cd78ccb2bb4f638.mp3',
                        cover: 'https://p1.music.126.net/5zs7IvmLv7KahY3BFzUmrg==/109951163635241613.jpg?param=300y300', // prettier-ignore
                        lrc: 'https://github.com/flipped-aurora/gf-vue-admin/tree/master/docs/东西-纳豆nado.lrc',
                    },
                    {
                        name: '响喜乱舞（Cover：MARiA）',
                        artist: '泠鸢yousa',
                        url: 'http://m10.music.126.net/20200815151334/8ffaf87dd4272d747c988c91e9168057/ymusic/3621/215a/86b5/fb61afc3dd541b41afbccbc4246026a5.flac',
                        cover: 'https://p1.music.126.net/AUGVPQ_rVrngDH9ocQrn3Q==/109951163613037822.jpg?param=300y300', // prettier-ignore
                        lrc: 'https://github.com/flipped-aurora/gf-vue-admin/tree/master/docs/响喜乱舞(泠鸢yousa).lrc',
                    },
                    {
                        name: '啵唧',
                        artist: 'Hanser',
                        url: 'http://m10.music.126.net/20200815151419/eae3a31b5e7eef9c52678ba97808f198/ymusic/47d8/5ede/5af4/41c2897b2f347bc2e28bce9a2341c05f.flac',
                        cover: 'https://p1.music.126.net/K0-IPcIQ9QFvA0jXTBqoWQ==/109951163636756693.jpg?param=300y300', // prettier-ignore
                        lrc: 'https://github.com/flipped-aurora/gf-vue-admin/tree/master/docs/啵唧-Hanser.lrc',
                    },
                ]
            }
        }
    }
</script>

<style scoped>

</style>
