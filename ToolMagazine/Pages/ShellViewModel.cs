using Stylet;


namespace ToolMagazine.Pages
{
    public class ShellViewModel : Screen
    {

        private string _name { get; set; }="wangkai";

        public string Name
        {
            get { return _name; }

            set
            {
                if (_name != value)
                {
                    this._name = value;
                    OnPropertyChanged(nameof(Name));
                }
              
            }
        }


        public void SayHello() => Name = "Hello " + Name;    // C#6的语法, 表达式方法

        public bool CanSayHello => !string.IsNullOrEmpty(Name);  // 同上

    }
}
