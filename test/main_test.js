import login_test from './login_test.js';
import load_test from './load_test.js';

export let options = {
  vus: 50,
  duration: '1m',
};

export default function () {
  login_test();
  load_test();
}