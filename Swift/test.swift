import UIKit

class ViewController: UIViewController {
    
    // ラベルとボタンを作成
    let label = UILabel()
    let button = UIButton(type: .system)
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        // ラベルの設定
        label.text = "Tap the button"
        label.textAlignment = .center
        label.frame = CGRect(x: 0, y: 200, width: view.frame.width, height: 50)
        view.addSubview(label)
        
        // ボタンの設定
        button.setTitle("Tap me", for: .normal)
        button.frame = CGRect(x: (view.frame.width - 100) / 2, y: 300, width: 100, height: 50)
        button.addTarget(self, action: #selector(buttonTapped), for: .touchUpInside)
        view.addSubview(button)
    }
    
    // ボタンがタップされたときのアクション
    @objc func buttonTapped() {
        label.text = "Button was tapped!"  // ラベルのテキストを変更
    }
}
