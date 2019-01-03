using Stylet;


namespace ToolMagazine.Pages
{
    public class ShellViewModel : Screen
    {

        private string email { get; set; }

        public string Email
        {
            get { return email; }

            set
            {
                if (email != value)
                {
                    this.email = value;
                    SetName(email);

                    //OnPropertyChanged(nameof(Email));
                    //NotifyPropertyChanged(Email);
                }

            }
        }


        private string _name { get; set; }

        public string Name
        {
            get { return _name; }

            set
            {              
                if (_name != value)
                {
                    this._name = value;             
                    //OnPropertyChanged(Name);
                    
                }

            }
        }

        private void SetName(string name)
        {

            if (!System.Text.RegularExpressions.Regex.IsMatch(name, @"\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*"))
            {
                this.Name = "请输入正确的邮箱格式";
            }
            else
            {
                this.Name = "";
            }
           
        } 

        public void SayHello() => Name = "Hello " + Name;    // C#6的语法, 表达式方法

        public bool CanSayHello => !string.IsNullOrEmpty(Name);  // 同上

    }
}
