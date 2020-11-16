import say from 'k6/x/say';

export default function () {
  console.log(say.hello("Daniel"));
}
