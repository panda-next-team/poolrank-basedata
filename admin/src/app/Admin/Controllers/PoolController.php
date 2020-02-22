<?php

namespace App\Admin\Controllers;

use App\Admin\Models\Pool;
use App\Http\Controllers\Controller;
use Encore\Admin\Controllers\HasResourceActions;
use Encore\Admin\Form;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;
use Encore\Admin\Show;
use App\Admin\Models\PowCoin;
use App\Admin\Models\PoolAddress;

class PoolController extends Controller
{
    use HasResourceActions;


    /**
     * Index interface.
     *
     * @param Content $content
     * @return Content
     */
    public function index(Content $content)
    {
        return $content
            ->header('List')
            ->description('Pool')
            ->body($this->grid());
    }

    /**
     * Show interface.
     *
     * @param mixed $id
     * @param Content $content
     * @return Content
     */
    public function show($id, Content $content)
    {
        return $content
            ->header('Detail')
            ->description('Pool')
            ->body($this->detail($id));
    }

    /**
     * Edit interface.
     *
     * @param mixed $id
     * @param Content $content
     * @return Content
     */
    public function edit($id, Content $content)
    {
        return $content
            ->header('Edit')
            ->description('Pool')
            ->body($this->form()->edit($id));
    }

    /**
     * Create interface.
     *
     * @param Content $content
     * @return Content
     */
    public function create(Content $content)
    {
        return $content
            ->header('Create')
            ->description('Pool')
            ->body($this->form());
    }

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new Pool);

        $grid->id('ID')->sortable();
        $grid->name('名称')->sortable();
        $grid->created_at('创建时间')->sortable();
        $grid->updated_at('更新时间')->sortable();

        $grid->status('状态')->display(function ($status) {
            return $status ? '<span style="color:darkgreen">启用</span>' : '<span style="color:darkred">停用</span>';
        });

        $grid->filter(function ($filter) {
            $filter->like('name', '名称');
        });

        return $grid;
    }

    /**
     * Make a show builder.
     *
     * @param mixed $id
     * @return Show
     */
    protected function detail($id)
    {
        $show = new Show(Pool::findOrFail($id));

        $show->id('ID');
        $show->name('名称');
        $show->field("icon", "icon")->as(function($icon) {
           return $icon ? sprintf("<img src=\"data:image/png;base64, %s\">", $icon) :"";
        })->unescape();
        $show->field("website_url", "官网")->link();
        $show->field('status', "状态")->as(function ($status) {
            return $status ? '<span style="color:darkgreen">启用</span>' : '<span style="color:#8b0000">停用</span>';
        })->unescape();

        $show->field('tags', "标记")->setRelation('tags')->as(function($tags){
            $data = [];
            foreach ($tags as $tag) {
                array_push($data, "<span style=\"color:mediumpurple\">". $tag['tag']."</span>");
            }
            return join(";", $data);
        })->unescape();

        $show->field('addresses', "地址")->setRelation("addresses")->as(function($addresses){
            $data = [];
            foreach ($addresses as $address) {
                array_push($data, "<span style=\"color:orange\">"."[".PoolAddress::$typeOptions[$address["type"]]. "]". $address['address']."</span>");
            }
            return join(";", $data);
        })->unescape();

        $show->created_at('创建时间');
        $show->updated_at('更新时间');
        return $show;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new Pool);
        $form->display('ID');

        $form->text("name", "名称")->rules('required');
        $form->url("website_url", "官方网站")->rules("required");
        //$form->image('icon', 'icon')->sequenceName()->removable()->rules('nullable');
        $form->textarea("icon", "ICON")->rules('required');


        $form->hasMany('tags', function (Form\NestedForm $form) {
            $form->text("tag", "标志")->placeholder(" ")->rules("required");
        });


        $coins = POWCoin::all();
        $coinOptions = [];
        foreach ($coins as $coin) {
            $coinOptions[$coin["id"]] = $coin["name"] . "-" . $coin["en_tag"];
        }

        $form->hasMany("addresses", function (Form\NestedForm $form) use ($coinOptions) {
            $form->select("coin_id", "币")->options($coinOptions)->rules("required");
            $form->select('type', "类型")->options(PoolAddress::$typeOptions)->rules("required");
            $form->text("address", "地址")->rules("required");
        });

        $form->select("status", "状态")->options([0 => '停用', 1 => '启用'])->default(1)->rules("required");
        $form->number("list_order", "排序")->default(0)->rules("required");
        return $form;
    }
}
