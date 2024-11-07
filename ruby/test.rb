# クラスの定義
class Person
    attr_accessor :name, :age
  
    # コンストラクタ
    def initialize(name, age)
      @name = name
      @age = age
    end
  
    # メソッドの定義
    def introduce
      puts "Hello, my name is #{@name} and I am #{@age} years old."
    end
  end
  
  # インスタンスを作成
  person = Person.new("Alice", 25)
  
  # メソッドを呼び出し
  person.introduce
  