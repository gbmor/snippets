extern crate crossbeam_channel;
use std::time;

pub fn cache_ticker() {
    let ticker = crossbeam_channel::tick(time::Duration::from_millis(60000));
    loop {
        select! {
            recv(ticker) -> _ => do_something(),
        }
    }
}

fn do_something() {
    unimplemented!();
}
