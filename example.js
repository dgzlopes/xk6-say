import hello from 'k6/x/say';

export default function () {
  console.log(hello.to("Daniel"));
}
