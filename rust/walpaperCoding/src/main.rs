use std::io;

fn main() {
    let nama = "Windy An-Nisa";

    println!("Aku mencintaimu, {}, karena Allah.", nama);

    println!("Apakah kamu ingin menjadi jodohku? (ya/tidak)");
    let mut jawaban = String::new();
    io::stdin()
        .read_line(&mut jawaban)
        .expect("Gagal membaca input");
    let jawaban = jawaban.trim().to_lowercase();

    if jawaban == "ya" {
        println!("Aku akan Membahagiakanmu,");
        println!("Aku akan Melidungimu dan Menjagamu,");
        println!("Aku akan Mengubah kamu menjadi lebih baik,");
        println!("Kalau ada Masalah Aku berharap kita Menyelesaikan nya Bersama-sama");
        println!("Semoga Allah jaga Hubungan kita hingga tua nanti");
    } else if jawaban == "tidak" {
        println!("Aku akan tetap mendoakanmu yang terbaik,");
        println!("Aku akan Menjaga mu hingga bertemu jodoh mu nanti");
    } else {
        println!("Aku juga akan meminta ke Allah agar kamu menjadi jodohku.");
    }
}
