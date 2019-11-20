<?php

namespace App\Admin\Controllers;

use App\Admin\Models\PowCoin;
use App\Admin\Models\Algorithm;
use App\Http\Controllers\Controller;
use Encore\Admin\Controllers\HasResourceActions;
use Encore\Admin\Form;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;
use Encore\Admin\Show;

class PowCoinController extends Controller
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
            ->description('POW Coin')
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
            ->description('POW Coin')
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
            ->description('POW Coin')
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
            ->description('POW Coin')
            ->body($this->form());
    }

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new PowCoin);
        $grid->id('ID')->sortable();
        $grid->name('名称')->sortable();
        $grid->column('en_name', "英文名称")->sortable();
        $grid->column('en_tag', "英文缩写")->sortable();
        $grid->column('algorithm.name', "算法名称");
        $grid->status('状态')->display(function ($status) {
            return $status ? '<span style="color:darkgreen">启用</span>' : '<span style="color:#8b0000">停用</span>';
        });
        $grid->created_at('创建时间');
        $grid->updated_at('更新时间');

        $grid->filter(function ($filter) {
            $filter->like('name', '名称');
            $filter->like('en_tag', '英文缩写');
            $filter->like('en_name', '英文名称');
            $filter->equal('status')->select(['0' => '停用', '1' => '启用']);
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
        $show = new Show(PowCoin::findOrFail($id));

        $show->id('ID');
        $show->field('name', "名称");
        $show->field('en_name', "英文名称");
        $show->field('en_tag', "英文缩写");
        $show->field('algorithm_id', "算法")->setRelation('algorithm')->as(function($data){
            return $data["name"];
        });
        $show->field('release_date', "发行日期");
        $show->field('max_supply', "最大供应总量");
        $show->field('block_time', "理论出块时间(秒)");
        $show->field("icon", "icon")->image();
        $show->field("github_url", "github")->link();
        $show->field("website_url", "官网")->link();
        $show->field("intro", "简介");
        $show->field('status', "状态")->as(function ($status) {
            return $status ? '<span style="color:darkgreen">启用</span>' : '<span style="color:#8b0000">停用</span>';
        })->unescape();
        $show->field('list_order', "自定义排序");
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

        $algorithms = Algorithm::all();
        $algorithmOptions = [];
        foreach ($algorithms as $algorithm) {
            $algorithmOptions[$algorithm["id"]] = $algorithm["name"];
        }

        $form = new Form(new PowCoin);
        $form->display('ID');
        $form->text("name", "名称")->placeholder(' ')->rules('required');
        $form->text("en_name", "英文名称")->placeholder(' ')->rules('required');
        $form->text("en_tag", "英文缩写")->placeholder('英文缩写建议大写')->rules('required');
        $form->select("algorithm_id", "算法")->options($algorithmOptions)->rules('required');
        $form->date('release_date', "发行日期")->placeholder(' ')->rules('nullable');
        $form->decimal("max_supply", "最大供应总量")->placeholder(' ')->rules('nullable');
        $form->number("block_time", "理论出块时间(秒)")->placeholder(' ')->rules('nullable');
        $form->image('icon', 'icon')->sequenceName()->removable()->rules('nullable');
        $form->url("github_url", "github")->rules('nullable');
        $form->url("website_url", "官网")->rules('nullable');
        $form->textarea('intro', '简介')->rules('nullable');
        $form->select("status", "状态")->options([0 => '停用', 1 => '启用'])->default(1)->rules("required");
        $form->number("list_order", "排序")->default(0)->rules("required");
        return $form;
    }
}
